package types

import . "goa.design/goa/v3/dsl"

// 返す値の構造、型を定義している
var Todo = ResultType("Todo", func() {
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("title", String, "Title")
		Attribute("is_done", Boolean, "IsDone")
	})
})
