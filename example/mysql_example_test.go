package example

import (
	"fmt"
	"go-rds-kit/config"
	"go-rds-kit/meta"
	"go-rds-kit/mysql"
	"go-rds-kit/types"
	"gorm.io/gorm"
	"testing"
)

const (
	UserTableName = "user"
)

type User struct {
	types.BaseModel
	Name string `json:"name" gorm:"column:username"`
	Age  int    `json:"age" gorm:"column:age"`
}

func (u User) String() string {

	return fmt.Sprintf("id:%d,name:%s,age:%d", u.ID, u.Name, u.Age)
}

type UserModel struct {
	mysql.TableOperate[User]
}

func NewUserModel(db *gorm.DB) *UserModel {

	return &UserModel{
		&mysql.Table[User]{
			DB:        db,
			TableName: UserTableName,
		},
	}
}

func Test_Sql(t *testing.T) {

	db := mysql.NewMysql(config.Mysql{
		Database: "database",
		Host:     "127.0.0.1",
		Passwd:   "11111",
		Port:     3306,
		User:     "22222",
		Logger:   true,
	}, nil)
	um := NewUserModel(db)

	user_info, err := um.Get(meta.GetOption{
		ID: 1,
	})
	fmt.Println(user_info, err)

	listOption := meta.ListOption{
		Page: &meta.Page{
			Page:     1,
			PageSize: 10,
		},
	}

	var (
		id   int    = 3
		name string = "name"
	)

	listOption.AddToCondition("id", &id)
	listOption.AddToCondition("name", &name)
	listOption.AddToCondition("age", nil) // 无效
	users, count, err := um.List(listOption)

	fmt.Println(users, count, err)

	var users2 []struct {
		Name    string
		Age     int
		Address string //  多表连接查询
	}

	aggregate, err := um.ListAggregate(meta.EmptyListOption, &users2) // 自定义结构查询
	fmt.Println(aggregate, users2, err)
}
