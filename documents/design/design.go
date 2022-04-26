package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("docs", func() {
	Title("API Documents")
	Description("HTTP service for API, a goa teaser")
	Server("docs", func() {
		Host("localhost", func() { URI("http://localhost:1323") })
	})
})

var _ = Service("guests", func() {
	Description("The calc service performs operations on numbers")

	Method("guests/:id", func() {
		Payload(func() {
			Attribute("id", Int, "Left operand")
			Required("id")
		})

		Result(Int)

		HTTP(func() {
			GET("/guests/{id}")
			Response(StatusOK)
		})
	})
	Method("guests", func() {
		Result(Int)

		HTTP(func() {
			GET("/guests")
			Response(StatusOK)
		})
	})
})
