package design

import . "goa.design/goa/v3/dsl"

// API全体を通しての説明とタイトルをつけられるもの、どのURLで待ち受けるかを指定する
var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service that manage todo.")
	Server("todo", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})
