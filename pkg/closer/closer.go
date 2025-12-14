package closer

const averageClosingsAtApp = 30

type closer interface {
	Close() error
}

var globalCloser = appCloser{
	make([]closer, 0, averageClosingsAtApp),
}

func Add(closeStruct closer) {
	globalCloser.closers = append(globalCloser.closers, closeStruct)
}

func AddFunc(f func() error) {
	globalCloser.closers = append(globalCloser.closers, newFuncCloser(f))
}
