# Goa で作る REST API サンプル

1. design/design.go で設定ファイルを作成
2. $ goa gen todo/design
   - gen/配下のファイルが自動生成される。（いじらない）
3. $ goa example todo/design
   - todo.go と cmd/配下が自動生成される
