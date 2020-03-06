package json

type JsonSchemaInfo struct {
	Schema interface{}
}

type JsonSchemaStore struct {
	Schemas map[string]JsonSchemaInfo
}

func NewJsonSchemaStore() JsonSchemaStore {
	return JsonSchemaStore{
		Schemas: make(map[string]JsonSchemaInfo),
	}
}

func (store *JsonSchemaStore) AddSchema(key string, info JsonSchemaInfo) {
	store.Schemas[key] = info
}

func (store *JsonSchemaStore) RemoveSchema(key string) {
	delete(store.Schemas, key)
}

func (store *JsonSchemaStore) Clear() {
	store.Schemas = make(map[string]JsonSchemaInfo)
}
