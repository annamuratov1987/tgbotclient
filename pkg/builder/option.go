package builder

type BuilderOption map[string]interface{}

func (o BuilderOption) SetOption(key string, value interface{}) {
	o[key] = value
}

func (o BuilderOption) GetOption(key string) interface{} {
	return o[key]
}
