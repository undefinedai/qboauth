package qboauth

func (d *Document) isFor(env Environment) bool {
	return (*d).environment == env
}
