package model

import (
	"errors"
	"gorm.io/gorm"
)

// 题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Age       int
	NickName  string
	PostCount int
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title        string `gorm:"unique"`
	Content      string
	UserID       uint
	CommentCount uint
	State        CommentState
	Comments     []Comment
}
type CommentState int

const (
	NO_COMMENT CommentState = iota + 1
	COMMENTED
)

func (p *Post) AfterCreate(db *gorm.DB) error {

	return db.Transaction(func(tx *gorm.DB) error {
		if p.UserID == 0 {
			return errors.New("无效的UserID")
		}

		result := db.Model(&User{}).
			Where("id = ?", p.UserID).
			Update("post_count", gorm.Expr("post_count+?", 1))

		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("用户不存在！！！")
		}
		return nil
	})
}

type Comment struct {
	gorm.Model
	Title     string `gorm:"unique"`
	Content   string
	StarCount uint
	PostID    uint
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Comment AfterDelete panic")
		}
	}()

	return db.Transaction(func(tx *gorm.DB) error {

		if c.PostID == 0 {
			return errors.New("无效的文章ID")
		}
		// 获取关联的文章
		var post Post
		if err := db.First(&post, c.PostID).Error; err != nil {
			return err
		}
		// 减少post 评论计数
		commentCount := post.CommentCount
		newCommentCount := commentCount - 1
		if newCommentCount < 0 {
			newCommentCount = 0
		}
		post.CommentCount = newCommentCount

		// 更新文章评论状态
		commentState := NO_COMMENT
		if newCommentCount > 0 {
			commentState = COMMENTED
		}
		post.State = commentState

		//更新post 数据
		return db.Model(&post).Select("comment_count", "state").Updates(&post).Error
	})
}
