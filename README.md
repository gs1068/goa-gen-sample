# Goa で作る REST API サンプル

1. design/design.go で設定ファイルを作成
2. $ goa gen todo/design
   - gen/配下のファイルが自動生成される。（いじらない）
3. $ goa example todo/design
   - todo.go と cmd/配下が自動生成される

参考
https://takakisan.com/golang-goa-beginner/

### SQL 準備

```
go get github.com/mattn/go-sqlite3

sqlite3 ./db.sqlite3

create table todos (
    id integer primary key autoincrement,
    title text not null,
    is_done boolean not null default false
);

insert into todos (title) values ('buy milk');

sqlite> select * from todos;
1|buy milk|0
```
