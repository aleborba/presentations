//START_1 OMIT
package main // HLmartini


import (
    "net/http" // HLmartini
    "database/sql"
    "github.com/coopernurse/gorp"
    "github.com/codegangsta/martini" // HLmartini
    "github.coopernursem/codegangsta/martini-contrib/render"
    "github.com/codegangsta/martini-contrib/binding"
    _ "github.com/go-sql-driver/mysql" // HLmartini
    "html/template" // HLmartini
    "log"
    "time"
)
//END_1 OMIT

//START_2 OMIT
func main() { // HLmartini

    m := martini.Classic() // HLmartini

    m.Use(render.Renderer(render.Options{
    m.Get("/", func(r render.Render) {
    m.Post("/", binding.Bind(Post{}), func(post Post, r render.Render) {
    m.Run() // HLmartini
}
// END_2 OMIT
