package app

type IGrep interface {
	IGetContent
	Grep(key , app string)
}

type IGetContent interface {
	GetContent(app string)
}