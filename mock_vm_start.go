package ovirtclient

func (m *mockClient) StartVM(id string, _ ...RetryStrategy) (VM, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.vms[id]; !ok {
		return nil, newError(ENotFound, "VM with ID %s not found", id)
	}
	m.vms[id].status = VMStatusUp
	return m.vms[id], nil
}