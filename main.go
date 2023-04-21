package main

func main() {
	r := NewRouter()

	r.Get("/hello", hello)

	r.Start(":4000")
}

func hello(r *Router) {
	println(r.Request.Header["Content-Type"][0])
	r.Response.Write([]byte("hello world from the handler"))
}
