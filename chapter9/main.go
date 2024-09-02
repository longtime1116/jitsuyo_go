package main

import (
	_ "database/sql"
)

func main() {
	// var err error
	// db, err := sql.Open()
	// rows, err := db.QueryContext()
	// defer rows.Close() // リソース解放
	// fmt.Println(err, rows)

	// var (
	// 	err error
	// 	db  *sql.DB
	// )
	// db, err = sql.Open()
	// tx, err := db.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer tx.Rollback() // deferでRollbackを登録
	// err = tx.QueryRow("UPDATE accounts SET balance = balance - 100 WHERE id = ?", fromAccountID).Err()
	// if err != nil {
	// 	return
	// }
	// err = tx.Commit() // ここでトランザクションを確定
	// if err != nil {
	// 	return
	// }
	// // 関数の最後でdeferによるRollbackが呼ばれる。トランザクションが確定していたら、このRollbackは呼ばれるものの意味は持たない
}
