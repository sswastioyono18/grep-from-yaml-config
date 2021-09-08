package app

type IGrep interface {
	IGetContent
	Grep(key , app string) bool
}

type IGetContent interface {
	GetContent(app string)
}