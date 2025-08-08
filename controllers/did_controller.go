package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mr-tron/base58"
	"github.com/satyakarki/did-api/domain"
	"github.com/satyakarki/did-api/services/crypto"
	"github.com/satyakarki/did-api/services/did"
)

// type DIDUsecase interface {
// 	CreateDID() (*domain.DIDDocument, string, error)
// }

type DIDController struct {
	Service *did.DIDService
}

func NewDIDController(service *did.DIDService) *DIDController {
	return &DIDController{Service: service}
}

func (u *DIDController) CreateDID(ctx echo.Context) error {
	// Generate Ed25519 key pair
	keyPair, err := crypto.GenerateEd25519KeyPair()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to generate key pair",
		})
	}

	// Simulate a did:key using UUID
	didID := "did:key:" + keyPair.PublicKey
	// Build DID Document
	doc := &domain.DIDDocument{
		Context: "https://www.w3.org/ns/did/v1",
		ID:      didID,
	}

	vm := domain.VerificationMethod{
		ID:              didID + "#keys-1",
		Type:            "Ed25519VerificationKey2020",
		Controller:      didID,
		PublicKeyBase58: keyPair.PublicKey,
	}

	doc.AddVerificationMethod(vm)

	// Save the DID to memory
	u.Service.SaveDID(*doc)

	// Respond with the DID Document (do NOT return the private key in API)
	//return ctx.JSON(http.StatusCreated, doc)

	return ctx.JSON(http.StatusOK, echo.Map{
		"did":        didID,
		"publicKey":  keyPair.PublicKey,
		"privateKey": keyPair.PrivateKey, // Show only for demo/testing
		"document":   doc,
	})
}

func (c *DIDController) ResolveDID(ctx echo.Context) error {
	id := ctx.Param("id")
	doc, err := c.Service.ResolveDID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "DID not found"})
	}
	return ctx.JSON(http.StatusOK, doc)
}

type AddKeyRequest struct {
	PublicKey string `json:"publicKey"`
	Type      string `json:"type"`
	KeyID     string `json:"keyId"`
}

func (c *DIDController) AddKeyToDIDHandler(ctx echo.Context) error {
	did := ctx.Param("id")

	var req AddKeyRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// Retrieve DID Document from service
	doc, err := c.Service.ResolveDID(did)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "DID not found"})
	}

	// Add key
	newKey := domain.VerificationMethod{
		ID:              did + "#" + req.KeyID,
		Type:            req.Type,
		Controller:      did,
		PublicKeyBase58: req.PublicKey,
	}

	doc.VerificationMethod = append(doc.VerificationMethod, newKey)
	doc.Authentication = append(doc.Authentication, newKey.ID)

	// Save updated DID
	c.Service.SaveDID(doc)

	return ctx.JSON(http.StatusOK, doc)
}

type SignRequest struct {
	Message string `json:"message"`
}

func (c *DIDController) SignMessageHandler(ctx echo.Context) error {
	// Simulate with private key (in production, securely retrieve it)
	message := ctx.FormValue("message")
	privateKey := ctx.FormValue("privateKey") // base58-encoded
	if message == "" || privateKey == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing message or privateKey"})
	}

	signature, err := crypto.SignMessageNew([]byte(message), privateKey)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Signing failed"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":   message,
		"signature": signature,
		//"publicKey": keyPair.PublicKey,
	})
}

type VerifyRequest struct {
	Message   string `json:"message"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

func (c *DIDController) VerifyMessageHandler(ctx echo.Context) error {
	message := ctx.FormValue("message")
	signature := ctx.FormValue("signature")
	publicKey := ctx.FormValue("publicKey")

	if message == "" || signature == "" || publicKey == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing fields"})
	}

	decodedPubKey, err := base58.Decode(publicKey)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid public key"})
	}

	decodedSig, err := base58.Decode(signature)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid signature"})
	}

	valid, err := crypto.VerifySignature([]byte(message), decodedSig, decodedPubKey)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Verification failed"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":   message,
		"signature": signature,
		"valid":     valid,
	})
}
