package postgres

import (
	pb "github.com/baxromumarov/template-service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

//create table users
func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	_, err := r.db.Query(`CREATE TABLE IF NOT EXISTS users (
		id varchar(255),
		first_name varchar(255),
		last_name varchar(255),
		email varchar(255),
		bio varchar(255),
		phoneNumbers text[],
	
		typeId varchar(255),
		Status varchar(255),
		createdAt varchar(255),
		updatedAt varchar(255),
		deletedAt varchar(255),
		FOREIGN KEY (address_id) REFERENCES addresses(id) )`)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//insert into users
func (r *userRepo) Insert(user *pb.User) (*pb.User, error) {
	var res = pb.User{}

	err := r.db.QueryRow(`INSERT INTO users (id, first_name, last_name, email, bio, phoneNumbers, typeId, Status, createdAt, updatedAt, deletedAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) Returning id,first_name,last_name`, user.Id, user.FirstName, user.LastName, user.Email,
		user.Bio, pq.Array(user.PhoneNumbers),
		user.TypeId, user.Status, user.CreatedAt, user.UpdatedAt, user.DeletedAt).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
	)
	if err != nil {
		return &pb.User{}, err
	}
	return &res, nil
}

//func (r *userRepo) Update(id, firstName, lastName *pb.User) (*pb.UserInfo, error) {
//	var res = pb.UserInfo{}
//
//	err := r.db.QueryRow(`UPDATE users SET first_name = $1, last_name = $2 where id = $3 returning id,first_name,last_name`, firstName, lastName, id).Scan(
//		&res.Id,
//		&res.FirstName,
//		&res.LastName,
//	)
//	if err != nil {
//		return &pb.UserInfo{}, err
//	}
//	return &res, nil
//}

func (r *userRepo) Delete(id *pb.ById) (*pb.UserInfo, error) {
	var res = pb.UserInfo{}

	_, err := r.db.Query(`DELETE FROM users where id = $1 returning id,
	first_name,last_name`,id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *userRepo) GetAll(user *pb.User) (*pb.User, error) {
	var res = pb.User{}

	err := r.db.QueryRow(`SELECT first_name, last_name, email, 
	bio, typeId, Status
	FROM users where id = $1  `, user.Id).Scan(
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Bio,
		&res.TypeId,
		&res.Status,
		
	)
	
	if err != nil {
		return &pb.User{}, err
	}
	return &res, nil
}
