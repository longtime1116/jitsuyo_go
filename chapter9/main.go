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

	// // プリペアードステートメントで効率的に大量のデータをinsert
	// // プリペアードステートメントとは？
	// //    プリペアードステートメントは、SQLクエリを事前に準備し、後で複数回にわたって異なるパラメータで実行する場合に非常に便利
	// // 		これにより、SQLインジェクション攻撃のリスクを減らし、また効率的にクエリを実行することができる。
	// //    具体的には、以下の二点で効率的になる。
	// //      DB側でクエリをparseし最適な実行計画をplanningする必要があるが、それが1回で済むのでCPU効率が良い
	// //      アプリケーション側から変数の部分のデータだけを送ることになるので、1回あたりのネットワークI/Oの負荷が減る。
	// //    命令としては、以下のような感じ。
	// //      PREPARE stmt FROM 'INSERT INTO users (name, age) VALUES (?, ?)';
	// //        EXECUTE stmt USING 'Alice', 30;
	// //        EXECUTE stmt USING 'Bob', 20;
	// //        ...
	// //
	// // バッチインサートとの使い分けはざっくり、以下
	// //   少量のデータを逐次挿入したい時、リアルタイム性が要求される時などは、プリペアードステートメントを検討。
	// //   より大規模なデータを一括で挿入したかったり、リアルタイムでなくバッチ処理が可能な場合は、バッチインサートを検討。
	// // なお、postgres の COPY のような組み込みの関数を使うという手段もある。
	// users := []struct {
	// 	UserId int
	// 	Name   string
	// }{
	// 	{1, "Ataro"},
	// 	{2, "Btaro"},
	// 	{3, "Ctaro"},
	// }
	// db, err := sql.Open()
	// ctx := context.Background()
	// tx, err := db.BeginTx(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// stmt, err := tx.PrepareContext(ctx, "INERT INTO users(user_id, user_name)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	// for _, u := range users {
	// 	if _, err := tx.ExecContext(ctx, u.UserId, u.Name); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// if err := tx.Commit(); err != nil {
	// 	log.Fatal(err)
	// }
}
