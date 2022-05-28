package repo

import (
	pb "github.com/baxromumarov/template-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	CreateAd(*pb.Address) (*pb.Address, error)
	Insert(*pb.User) (*pb.User, error)
	InsertAd(*pb.Address) (*pb.Address, error)
	//Update(id, firstName, lastName *pb.User) (*pb.UserInfo, error)
	Delete(id *pb.ById) (*pb.UserInfo, error)
	GetAll(*pb.ById) (*pb.User, error)
}

/*

need the method: GetAll(id *pb.ById) (*pb.User, error)
have the method: GetAll(user *pb.User) (*pb.User, error)
*/
