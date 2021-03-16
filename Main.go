package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func handleFun() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}

func main() {
	handleFun()

}


























//package main
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//)

//type User struct {
//	Name      string
//	Age       uint16
//	Money     int
//	AvgGrades float64
//	Happiness float64
//	Hobbies   []string
//}
//
//func (u User) getAllInfo() string {
//	return fmt.Sprintf("User name is: %s and he is %d and he has money  equal: %d", u.Name, u.Age, u.Money)
//}
//
//func (u *User) setName(newName string) {
//	u.Name = newName
//}
//
//func homePage(w http.ResponseWriter, r *http.Request) {
//	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Skate, Dance, Football"}}
//	tmpl, _ := template.ParseFiles("templates/home_page.html")
//	tmpl.Execute(w, bob)
//
//}
//
//func contactsPage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Contacts")
//}
//
//func handleRequest() {
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/contacts/", contactsPage)
//	http.ListenAndServe(":8888", nil)
//}







//database




//type User struct {
//	Name string `json:"name"` //json output format
//	Age uint16 `json:"age"`
//}
//
//
//func main() {
//	//handleRequest()
//
//	db, err := sql.Open("mysql", "root:72305@tcp(127.0.0.1:3306)/golang")
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//
//	// insert data
//	//insert, err := db.Query("insert into users (name, age) VALUES ('Bob', 30)")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer insert.Close()
//
//
//	res, err := db.Query("select name, age from users")
//	if err != nil {
//		panic(err)
//	}
//
//	for res.Next() {
//		var user User
//		err = res.Scan(&user.Name, &user.Age) //достаем и сохраняем в объект
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("User: %s with age %d \n", user.Name, user.Age)
//	}
//
//}
