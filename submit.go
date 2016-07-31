package main

import (
        "fmt"
        "os"
        "io/ioutil"
        "bufio"
        "strings"
        "database/sql"
        _ "github.com/lib/pq"
        "github.com/stvp/go-toml-config"
)

type Post struct {
      title string
      subtitle string
      content string
}

func loadPost(title string) (string, error) {
      filename := "texts/" + title + ".md"
      body, err := ioutil.ReadFile(filename)
      checkErr(err)
      return string(body), nil
}

func storePost(post Post, url string) {
  db, err := sql.Open("postgres", url)
  checkErr(err)
  fmt.Println("Inserting values into " + url)
  id, err := db.Exec("INSERT INTO posts(title,subtitle,content,\"createdAt\",\"updatedAt\") VALUES($1,$2,$3,now(), now())",
                    post.title, post.subtitle, post.content)
  checkErr(err)
  fmt.Println(id)
  fmt.Println("Success")
}

func main() {
  var url string
  config.StringVar(&url, "url", "postgres://test:test@localhost:5432/test")
  config.Parse("./submit.conf")
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter title: ")
  title, _ := reader.ReadString('\n')
  fmt.Print("Enter subtitle: ")
  subtitle, _ := reader.ReadString('\n')
  fmt.Print("Enter filename: ")
  filename, _ := reader.ReadString('\n')
  readContent, err := loadPost(strings.TrimSpace(filename))
  checkErr(err)
  newPost := Post{title: strings.TrimSpace(title), subtitle: strings.TrimSpace(subtitle), content: readContent}
  storePost(newPost, url)
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
