package repo

import (
	pb "github.com/baxromumarov/template-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
}
