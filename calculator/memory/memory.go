package memory

type memory struct {
	bytes int
}

func (m *memory) Setmemory(val int) {
	m.bytes = val
	return
}
func (m *memory) AddtoMemory(value int) int {
	m.bytes += value
	return m.bytes
}
func (m *memory) SubfromMemory(value int) int {
	m.bytes -= value
	return m.bytes
}
func (m *memory) ResetMemory() {
	m.bytes = 0
}
func Factory() *memory {
	return &memory{}
}
