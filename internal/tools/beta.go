package tools

type (
	Beta interface {
		Beta()
	}

	beta struct{}
)

func NewBeta() Beta {
	return &beta{}
}

func (b *beta) Beta() {
	// Do something beta.
}
