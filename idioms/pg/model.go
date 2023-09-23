package pg

type IQuery interface {
	ToSQL() string
}
