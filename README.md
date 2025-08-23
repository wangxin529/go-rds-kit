# GO-RDS-KIT

用于 mysql 和 redis的统一 封装

## 快速开始

### 1. 引入依赖
```shell
go get https://github.com/wangxin529/go-rds-kit
```

### 2. 配置

### mysql 配置

```go

type Mysql struct {
	User     string `json:"user"`  
	Passwd   string `json:"passwd"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Logger   bool   `json:"logger"`   // sql 日志打印
	DNS      string `json:"dns"`     // mysql 连接地址  
}


当填充DNS时，其他参数无效
```

### 缓存配置
```go

type Memory struct {
Type  CacheType `json:"type"` // 缓存类型 支持 memory 和 redis 两种
Redis *Redis    `json:"redis,optional"`
}
type Redis struct {
Addr             []string `json:"addr"` //redis 地址 一个为单点Redis 一个为集群Redis 多个地址为哨兵 Redis 
Password         string   `json:"password"`
MasterName       string   `json:"masterName"` // 哨兵模式参数
SentinelPassword string   `json:"sentinelPassword"` // 哨兵模式参数
DB               int      `json:"db"`
}

```

## 3.使用说明 

### 3.1. mysql 使用
```go

const (
	UserTableName = "user"
)

type User struct {
    types.BaseModel
    Name string `json:"name" gorm:"column:username"`
    Age  int    `json:"age" gorm:"column:age"`
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

func main(){

	db := mysql.NewMysql(config.Mysql{}, nil)
	um := NewUserModel(db)
	user := &User{
		Name: "test",
		Age:  18,
	}
	err := um.Create(user)

	fmt.Println(err)

	user_info, err := um.Get(meta.GetOption{
		ID: 1,
	})
	fmt.Println(user_info, err)

	err = um.Update(meta.UpdateOption{
		ID: user.ID,
		Data: User{
			Name: "test1",
			Age:  18,
		},
	})
	fmt.Println(err)

	err = um.Delete(meta.DeleteOption{
		ID: user.ID,
	})
	fmt.Println(err)

	users, count, err := um.List(meta.ListOption{
		Page: &meta.Page{
			Page:     1,
			PageSize: 10,
		},
	})

	fmt.Println(users, count, err)

	var users2 []struct {
		Name    string
		Age     int
		Address string //  多表连接查询
	}

	aggregate, err := um.ListAggregate(meta.EmptyListOption, &users2) // 自定义结构查询
	fmt.Println(aggregate, users2, err)
}


```

### 3.2 缓存使用
```go
func NewCache(conf *config.Memory) cache.Cache {

	return cache.NewMemory(conf)

}
func Test_Redis(t *testing.T) {

	//1. 内存
	memory := NewCache(&config.Memory{
		Type: config.CacheMemoryType,
	})
	memory.Set("key", "value", -1)

	//2. 单点 redis
	redis := NewCache(&config.Memory{
		Type: config.CacheRedisType,
		Redis: &config.Redis{
			Addr:     []string{"127.0.0.1:6379"},
			DB:       0,
			Password: "",
		},
	})
	redis.Set("key", "value", -1)

	// 3. 哨兵 redis
	redis = NewCache(&config.Memory{
		Type: config.CacheRedisType,
		Redis: &config.Redis{
			MasterName:       "mymaster",
			SentinelPassword: "",
			Addr:             []string{"127.0.0.1:26379","127.0.0.1:6379"},
			Password:         "",
			DB:               0,
		},
	})
	redis.Set("key", "value", -1)

}
```
