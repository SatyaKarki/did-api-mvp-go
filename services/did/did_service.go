package did

import (
	"errors"
	"sync"

	"github.com/satyakarki/did-api/domain"
)

type DIDService struct {
	store map[string]domain.DIDDocument
	lock  sync.RWMutex
}

func NewDIDService() *DIDService {
	return &DIDService{
		store: make(map[string]domain.DIDDocument),
	}
}

func (s *DIDService) SaveDID(doc domain.DIDDocument) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[doc.ID] = doc
}

func (s *DIDService) SaveDIDDocument(doc *domain.DIDDocument) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[doc.ID] = *doc
}

func (s *DIDService) ResolveDID(id string) (domain.DIDDocument, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	doc, ok := s.store[id]
	if !ok {
		return domain.DIDDocument{}, errors.New("DID not found")
	}
	return doc, nil
}
