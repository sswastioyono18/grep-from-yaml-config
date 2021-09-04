package main

type IGrep interface {
	IGetContent
}

type IGetContent interface {
	GetContent() (yamlData map[string]interface{})
	Grep(key string)
}
