package cmd

type IGrep interface {
	IGetContent
	Grep(key string)
}

type IGetContent interface {
	GetContent() (yamlData map[string]interface{})
}
