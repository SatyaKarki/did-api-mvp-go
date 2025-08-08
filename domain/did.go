package domain

// DIDDocument represents a basic W3C-compliant DID Document
type DIDDocument struct {
	Context            string               `json:"@context"`
	ID                 string               `json:"id"`
	VerificationMethod []VerificationMethod `json:"verificationMethod"`
	Authentication     []string             `json:"authentication"`
	Service            []ServiceEndpoint    `json:"service,omitempty"`
}

// VerificationMethod holds key and verification details
type VerificationMethod struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Controller      string `json:"controller"`
	PublicKeyBase58 string `json:"publicKeyBase58"`
}

// ServiceEndpoint represents a service tied to the DID
type ServiceEndpoint struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

//helper
func (doc *DIDDocument) AddVerificationMethod(vm VerificationMethod) {
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	doc.Authentication = append(doc.Authentication, vm.ID)
}
