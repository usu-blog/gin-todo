package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Todo Type
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DBに接続
func dbOpen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbOpen）")
	}
	return db
}

// Init 初期化
func Init() {
	db := dbOpen()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// Insert 追加
func Insert(text string, status string) {
	db := dbOpen()
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// Update 更新
func Update(id int, text string, status string) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// Delete db delete
func Delete(id int) {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// GetAll 全取得
func GetAll() []Todo {
	db := dbOpen()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// GetOne 一つ取得
func GetOne(id int) Todo {
	db := dbOpen()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
