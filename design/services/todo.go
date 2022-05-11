package services

import (
	"todo/design/types"

	. "goa.design/goa/v3/dsl"
)

// 具体的なAPIメソッド達を定義する
var _ = Service("todo", func() {
	Description("Service that manage todo.")
	Method("hello", func() {
		// どんな値を受け取るか
		Payload(func() {
			// ここでは、"name"という名前で、文字列を受け取る
			Attribute("name", String, "Name")
			// Requireとすると必須項目にできる
			Required("name")
		})
		// どんな値を返すか
		// Resultで型を指定
		Result(String)
		HTTP(func() {
			GET("/hello/{name}")
			Response(StatusOK)
		})
	})
})

var _ = Service("todo", func() {
	// (略)
	Method("show", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Required("id")
		})
		Result(types.Todo)
		HTTP(func() {
			GET("/todo/{id}")
			Response(StatusOK)
		})
	})
})
