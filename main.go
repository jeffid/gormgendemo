package main

import (
	"context"
	"fmt"
	"log"

	"gormgendemo/dal/model"
	"gormgendemo/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:123456@(localhost:3306)/gormgendemo?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	query.SetDefault(db)
	q := query.Q
	ctx := context.Background()
	//qc := q.WithContext(ctx)

	// 增
	insert(ctx, q)
	// 删
	del(ctx, q)
	// 改
	update(ctx, q)
	// 查
	find(ctx, q)

	fmt.Println("Done!")
}

func insert(ctx context.Context, q *query.Query) {
	qc := q.WithContext(ctx)

	// 插入数据
	users := []*model.User{
		{
			Name: "张三",
			Age:  30,
			Address: []model.Address{
				{Province: "广东", City: "广州"},
				{Province: "广东", City: "深圳"},
			},
		},
		{
			Name: "李四",
			Age:  40,
			Address: []model.Address{
				{Province: "广东", City: "广州"},
				{Province: "广东", City: "深圳"},
			},
		},
		{
			Name: "王五",
			Age:  50,
		},
	}
	hobby := []*model.Hobby{
		{Name: "看电影"}, {Name: "看书"}, {Name: "跑步"},
	}

	err := qc.User.Create(users...)
	if err != nil {
		log.Fatal(err)
	}
	err = qc.Hobby.Create(hobby...)
	if err != nil {
		log.Fatal(err)
	}
}

func del(ctx context.Context, q *query.Query) {
	qc := q.WithContext(ctx)

	qc.User.Where(q.User.Name.Eq("张三")).Delete()       // 软删
	qc.Address.Where(q.Address.City.Eq("深圳")).Delete() // 软删
	qc.Hobby.Where(q.Hobby.Name.Eq("看电影")).Delete()    // 非软删
}

func update(ctx context.Context, q *query.Query) {
	qc := q.WithContext(ctx)

	qc.User.Where(q.User.Name.Eq("李四")).UpdateSimple(q.User.Age.Add(1))
}

func find(ctx context.Context, q *query.Query) {
	qc := q.WithContext(ctx)

	// First、Take、Last 方法, 如果没有找到记录则返回错误 ErrRecordNotFound。

	_, err := qc.User.Preload(q.User.Address).Where(q.User.Name.Eq("张三")).First()
	if err != nil {
		fmt.Printf("%+v \n", err)
		//record not found
	}
	user, err := qc.User.Preload(q.User.Address).Where(q.User.Name.Eq("李四")).First()
	if err == nil {
		fmt.Printf("%+v \n", *user)
		//{ID:2 Name:李四 Age:41 Balance:0 UpdatedAt:2022-09-16 15:13:39 +0800 CST CreatedAt:2022-09-16 15:13:39 +0800 CST
		//DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}
		//Address:[{ID:3 UID:2rovince:广东 City:广州 UpdateTime:1663312418 CreateTime:1663312418 DeleteTime:0
		//User:{ID:0 Name: Age:0 Balance:0 UpdatedAt:0001-01-01 00:00:00 +0000 UTC CreatedAt:0001-01-01 00:00:00 +0000 UTC
		//CreatedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false} Address:[]}}]}
	}

	addr, err := qc.Address.Preload(q.Address.User).Offset(1).First()
	if err == nil {
		fmt.Printf("%+v \n", *addr)
		//{ID:3 UID:2 Province:广东 City:广州 UpdateTime:1663312418 CreateTime:1663312418 DeleteTime:0
		//User:{ID:2 Name:李四 Age:41 Balance:0 UpdatedAt:2022-09-16 15:13:39 +0800 CST CreatedAt:2022-09-16 39 +0800 CST
		//DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false} Address:[]}}
	}

	_, err = qc.Hobby.Where(q.Hobby.Name.Eq("看电影")).First()
	if err != nil {
		fmt.Printf("%+v \n", err)
		//record not found
	}
	hb, err := qc.Hobby.Where(q.Hobby.Name.Eq("看书")).First()
	if err == nil {
		fmt.Printf("%+v \n", *hb)
		//{ID:2 Name:看书 UpdatedAt:1663312418 CreatedAt:1663312418 DeletedAt:0}
	}

}
