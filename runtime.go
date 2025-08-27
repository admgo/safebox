package safebox

type Runtime interface {
	Exec(c Command) Result
}
