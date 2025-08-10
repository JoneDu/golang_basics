package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"task1/task3/gorm/model"
	"time"
)

//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

var db *gorm.DB
var err error

func init() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			//ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)

	// 初始化DB
	dsn := "root:123456@tcp(127.0.0.1:3306)/school_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	// 初始化数据表(自动迁移数据表)
	//db.AutoMigrate(&model.Post{}, &model.User{}, &model.Comment{})

	fmt.Println("连接成功！")
}

func CreateTable() {
	// init func will createTable
}

// region 数据插入测试
func InsertData() {
	user := model.User{
		Name:     "小一",
		Age:      10,
		NickName: "一号选手",
		Posts: []model.Post{
			{
				Title:   "第一篇文章",
				Content: "文章内容",
				Comments: []model.Comment{
					{
						Title:     "第一篇文章的第一个评论",
						Content:   "评论内容",
						StarCount: 22,
					},
				},
			},
		},
	}

	db.Create(&user)
}

func InsertData2() {
	user := model.User{
		Name:     "小二",
		Age:      20,
		NickName: "二号选手",
	}
	db.Create(&user)

	// 显示创建二号选手的文章两条
	post1 := model.Post{
		Title:   "这是第二篇文章",
		Content: "文章内容",
		UserID:  user.ID,
	}
	post2 := model.Post{
		Title:   "这是第三篇文章",
		Content: "文章内容",
		UserID:  user.ID,
	}
	posts := []*model.Post{&post1, &post2}
	db.Model(&user).Association("Posts").Append(posts)

	// 给文章创建两条评论， post1
	comment1 := model.Comment{
		Title:     "第二个评论",
		Content:   "评论内容",
		StarCount: 0,
		PostID:    post1.ID,
	}
	db.Model(&post1).Association("Comments").Append([]model.Comment{comment1})

	//comment2 := model.Comment{
	//	Title:     "第三个评论",
	//	Content:   "评论内容",
	//	StarCount: 22,
	//	PostID:    post2.ID,
	//}
	//db.Create(&comment2)

}

func ConflictInsert() {
	user := model.User{
		Name:     "小四",
		Age:      40,
		NickName: "四号选手",
	}
	db.Create(&user)

	post1 := model.Post{
		Title:   "这是第四篇文章",
		Content: "文章内容测试冲突",
		UserID:  user.ID,
	}
	db.Model(&user).Association("Posts").Append(&post1)
	comment1 := model.Comment{
		Title:     "第二个评论",
		Content:   "评论title冲突",
		StarCount: 1000,
		PostID:    post1.ID,
	}

	// 指定conflict clauses
	appendCommentWithConflictUpdate(db, &post1, &comment1, []string{"Content", "StarCount"})
}

func appendCommentWithConflictUpdate(db *gorm.DB, post *model.Post, comment *model.Comment, updateFields []string) error {
	// 设置冲突条款
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "Title"}},
		DoUpdates: clause.AssignmentColumns(updateFields),
	})

	// 关联保存Comments，但是绕过Post的关联保存，并自定义Post 更新
	err2 := db.Session(&gorm.Session{}).Model(post).Association("Comments").Append(comment)

	// 更新post
	post.Content = post.Content + "===我被更新了"
	post.UpdatedAt = time.Now()
	db.Model(post).Select("Content", "UpdatedAt").Updates(post)

	return err2
}

func ConflictInsertComment() {
	var post model.Post
	db.First(&post, 6)

	var comment model.Comment
	db.First(&comment, "post_id=?", post.ID)
	comment.Title = "This is a title333"
	comment.Content = "这是55555"
	comment.UpdatedAt = time.Now()

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "title"}},
		DoUpdates: clause.AssignmentColumns([]string{"content", "updated_at"}),
	}).Omit("id").Create(&comment)

}

//endregion

// 题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

func QueryData() {
	var user model.User
	db.Preload("Posts.Comments").First(&user, 4)
	fmt.Printf("user: %+v\n", user)

	// 评论数 最多的文章
	var posts []model.Post
	db.Preload("Comments").
		Select("posts.*,count(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Find(&posts)
	fmt.Printf("posts: %+v\n", posts)
}

//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func CreatPost() {
	post := model.Post{
		Title:   "文章标题22-0-",
		Content: "文222章内容",
		UserID:  4,
	}
	db.Create(&post)
}

func DeleteComment() {
	var comment model.Comment
	db.First(&comment, 7)
	db.Delete(&comment, 7)
}
