package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func loadDb(path string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := createTableIfNotExists(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTableIfNotExists(db *sql.DB) error {
	// Create table if not exists
	file, err := os.Open(DbSchemaPath)
	if err != nil {
		return err
	}
	defer file.Close()
	var schema string
	if _, err := file.Read([]byte(schema)); err != nil {
		return err
	}
	_, err = db.Exec(schema) // CREATE TABLE IF NOT EXIST ...;
	return err
}

func scanItem(rows *sql.Rows) (*Item, error) {
	var item Item
	if err := rows.Scan(&item.Id, &item.Name, &item.CategoryId, &item.ImageName); err != nil {
		return nil, err
	}
	return &item, nil
}

func scanCategory(rows *sql.Rows) (*Category, error) {
	var category Category
	if err := rows.Scan(&category.Id, &category.Name); err != nil {
		return nil, err
	}
	return &category, nil
}

func scanJoinedItem(rows *sql.Rows) (*JoinedItem, error) {
	var joined_item JoinedItem
	if err := rows.Scan(&joined_item.Id, &joined_item.Name, &joined_item.CategoryName, &joined_item.ImageName); err != nil {
		return nil, err
	}
	return &joined_item, nil
}

func scanJoinedItems(rows *sql.Rows) (*JoinedItems, error) {
	var joined_items JoinedItems
	for rows.Next() {
		joined_item, err := scanJoinedItem(rows)
		if err != nil {
			return nil, err
		}
		joined_items.Items = append(joined_items.Items, *joined_item)
	}
	return &joined_items, nil
}

func loadItemById(db *sql.DB, id int) (*Item, error) {
	// Load item from db by id
	rows, err := db.Query("SELECT * FROM items WHERE items.id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanItem(rows)
}

func loadCategoryById(db *sql.DB, id int) (*Category, error) {
	// Load category from db by id
	rows, err := db.Query("SELECT * FROM categories WHERE categories.id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanCategory(rows)
}

func insertItem(db *sql.DB, item Item) error {
	// Save new items to database
	_, err := db.Exec("INSERT INTO items (name, category_id, image_name) VALUES (?, ?, ?)", item.Name, item.CategoryId, item.ImageName)
	return err
}

func insertCategory(db *sql.DB, category Category) error {
	// Save new category to database
	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	return err
}

func searchItemsByKeyword(db *sql.DB, keyword string) (*JoinedItems, error) {
	query := JoinAllQuery + " AND items.name LIKE CONCAT('%', ?, '%')"
	fmt.Println(query)
	rows, err := db.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanJoinedItems(rows)
}

func joinItemAndCategory(db *sql.DB, item Item) (*JoinedItem, error) {
	category, err := loadCategoryById(db, item.CategoryId)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, nil
	}

	joined_item := JoinedItem{Id: item.Id, Name: item.Name, ImageName: item.ImageName, CategoryName: category.Name}
	return &joined_item, nil
}

func joinAll(db *sql.DB) (*JoinedItems, error) {
	// Join category name to items
	rows, err := db.Query(JoinAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanJoinedItems(rows)
}
