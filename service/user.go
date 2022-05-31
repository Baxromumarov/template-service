package service

import (
	"context"

	"time"

	pb "github.com/baxromumarov/user-service/genproto"
	l "github.com/baxromumarov/user-service/pkg/logger"
	"github.com/baxromumarov/user-service/storage"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// func (s *UserService) Update(ctx context.Context, id *pb.ById) (*pb.UserInfo, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (s *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.UserInfo, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (s *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error while creating user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating user")
	}
	return user, nil
}

func (s *UserService) CreateAd(ctx context.Context, cad *pb.Address) (*pb.Address, error) {

	cred, err := s.storage.User().CreateAd(cad)
	if err != nil {
		s.logger.Error("Error while creating address", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating address")
	}
	return cred, nil
}

func (s *UserService) Insert(ctx context.Context, req1 *pb.User) (*pb.User, error) {
	id, err := uuid.NewV4()
	crtime := time.Now()

	if err != nil {
		s.logger.Error("Error while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while generating uuid")
	}
	req1.Id = id.String()
	req1.CreatedAt = crtime.UTC().Format(time.RFC3339)
	user, err := s.storage.User().Insert(req1)
	if err != nil {
		s.logger.Error("Error while inserting user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting user")
	}
	return user, nil

}
func (s *UserService) InsertAd(ctx context.Context, add *pb.Address) (*pb.Address, error) {
	idd, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("Error while inserting address", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting address")
	}
	add.Id = idd.String()
	address, err := s.storage.User().InsertAd(add)
	if err != nil {
		s.logger.Error("Error while inserting address", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting address")
	}
	return address, nil
}

//func (s *UserService) Update(ctx context.Context, id, firstName, lastName *pb.User) (*pb.UserInfo, error) {
//	user, err := s.storage.User().Update(id, firstName, lastName)
//	if err != nil {
//		s.logger.Error("Error while updating user", l.Error(err))
//		return nil, status.Error(codes.Internal, "Error while updating user")
//	}
//	return user, nil
//}

func (s *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.UserInfo, error) {
	user, err := s.storage.User().Delete(id)
	if err != nil {
		s.logger.Error("Error while deleting user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while deleting user")
	}
	return user, nil
}

func (s *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
	user, err := s.storage.User().GetById(id)
	if err != nil {
		s.logger.Error("Error while getting all users", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users")
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pb.Empty) (*pb.UserResp, error) {
	//var users []*pb.User
	users, err := s.storage.User().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all users", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users")
	}

	return &pb.UserResp{
		User: users,
	},nil

}
