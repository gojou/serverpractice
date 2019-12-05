package server

import "context"

//Server interface
type Server interface {
	Get(ctx context.Context, id int64) (interface{}, err)
}
