package ovirtclient

import (
	"sync"

	"github.com/google/uuid"
)

// MockClient provides in-memory client functions, and additionally provides the ability to inject
// information.
type MockClient interface {
	Client

	// GenerateUUID generates a UUID for testing purposes.
	GenerateUUID() string
}

type mockClient struct {
	url            string
	lock           *sync.Mutex
	storageDomains map[string]storageDomain
	disks          map[string]disk
	clusters       map[string]cluster
	hosts          map[string]host
	templates      map[string]template
}

func (m *mockClient) GetURL() string {
	return m.url
}

func (m *mockClient) GenerateUUID() string {
	return uuid.NewString()
}