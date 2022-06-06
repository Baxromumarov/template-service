package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/baxromumarov/my-services/user-service/genproto"
	l "github.com/baxromumarov/my-services/user-service/pkg/logger"
	cl "github.com/baxromumarov/my-services/user-service/service/grpc_client"
	"github.com/baxromumarov/my-services/user-service/storage"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client cl.GrpcClientI
}


//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger,client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client: client,
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
	if req1.Post != nil {
		for _, post := range req1.Post {
			post.UserId = user.Id
			createdPost , err := s.client.PostSevice().CreatePost(context.Background(), post)
			if err != nil {
				s.logger.Error("Error while inserting post", l.Error(err))
				return nil, status.Error(codes.Internal, "Error while inserting post")
			}
			fmt.Println(createdPost)
		}
	
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
	users, err := s.storage.User().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all users1", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users1")
	}

	for _, user := range users {
		posts, err := s.client.PostSevice().GetAllUserPosts(
			ctx, 
			&pb.ByUserIdPost{UserId: user.Id,
				},
			) 
	
	if err != nil {
		s.logger.Error("Error while getting all users", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users")
	}
	user.Post = posts.Posts

	}
	return &pb.UserResp{
		User: users,
	},err

}

func (s *UserService) GetAllUserPosts(ctx context.Context, req *pb.ByUserIdPost) (*pb.GetUserPosts, error) {
	address, err := s.client.PostSevice().GetAllUserPosts(ctx, req)
	if err != nil {
		return nil, err
	}
	return address, nil
}