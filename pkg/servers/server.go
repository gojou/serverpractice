package server

type Server interface {
	Get(ctx context.Context, id int64) (interface{}, err)
}
