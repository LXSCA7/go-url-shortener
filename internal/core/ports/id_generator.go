package ports

type IDGenerator interface {
	Generate() int64
}
