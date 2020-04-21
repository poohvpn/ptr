package ptr

func (p *Time) Equal(other *Time) bool {
	return p.Value().Equal(other.Value())
}
