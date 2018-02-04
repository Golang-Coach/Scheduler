package interfaces

type IContext interface {
	Query(key string)
	DefaultQuery(key, defaultValue string)
}
