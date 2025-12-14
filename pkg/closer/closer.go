package closer

const averageClosingsAtApp = 30

type closer interface {
	Close() error
}

var globalCloser = newAppCloser(averageClosingsAtApp)

func Add(closeStruct closer) {
	globalCloser.addCloser(closeStruct)
}

func AddFunc(f func() error) {
	globalCloser.addCloser(newFuncCloser(f))
}
