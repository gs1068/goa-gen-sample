package services

import (
	"todo/design/types"

	. "goa.design/goa/v3/dsl"
)

// 具体的なAPIメソッド達を定義する
// おそらくattributeで指定しているものはパラメータとjsonどちらで渡しても問題なし
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

var _ = Service("todo", func() {
	// (略)
	Method("create", func() {
		Payload(func() {
			Attribute("title", String, "Title")
			Required("title")
		})
		Result(String)
		HTTP(func() {
			POST("/todo")
			Response(StatusOK)
		})
	})
})

var _ = Service("todo", func() {
	// (略)
	Method("update", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Attribute("is_done", Boolean, "IsDone")
			Required("id")
			Required("is_done")
		})
		Result(String)
		HTTP(func() {
			PUT("/todo/{id}")
			Response(StatusOK)
		})
	})
})

var _ = Service("todo", func() {
	// (略)
	Method("delete", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Required("id")
		})
		Result(String)
		HTTP(func() {
			DELETE("/todo/{id}/delete")
			Response(StatusOK)
		})
	})
})
