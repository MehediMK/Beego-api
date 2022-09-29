package controllers

import (
	//"encoding/json"

	"fmt"

	"regexp"
	//"github.com/beego/beego/v2/client/httplib"
	//"github.com/beego/beego/v2/core/validation"

	//"beego-api/models"
	//"net/http"
	//"html/template"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

// type User struct {
//     Id    int         `form:"-"`
//     Name  interface{} `form:"username"`
//     Age   int         `form:"age"`
//     Email string
// }

type User struct {
    FirstName      string  
	LastName       string  
	Email          string  
	Phone          string  
	Password       string  
	DateOfBirth    string  
	
}

// Operations about Users
type UserController struct {
	beego.Controller
}





// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	//FirstName  := `form:"fname"`
	u.TplName = "form.tpl"
	u.Data["fname"] = u.GetString("fname")
	u.Data["lname"] = u.GetString("lname")
	u.Data["phone"] = u.GetString("phone")
	u.Data["pass"] = u.GetString("pass")
	u.Data["email"] = u.GetString("email")  
	u.Data["dob"] = u.GetString("dob")
	

	//Phone Number validation
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	// End

	str := fmt.Sprintf("%v", u.Data["phone"])
	fmt.Println(re.MatchString(str)) 
	fmt.Println(u.Data["phone"]) 

	flash := web.NewFlash()
	flash.Error("Settings invalid!")
	u.Redirect("http://localhost:8081/form", 302)
	//u.ParseForm(&setting)
	// req := httplib.Get("http://beego.vip/")
	// fmt.Println(req) 
	//t, err := template.ParseFile("template/search.html")
	//http.Redirect(w, r, fmt.Sprintf("search?q=%s", query), 302)
	//Redirect("/setting", 302)
	// flash := web.NewFlash()
    // setting := Settings{}
    // valid := Validation{}
    // c.ParseForm(&setting)
	// if b, err := valid.Valid(setting); err != nil {
    //     flash.Error("Settings invalid!")
    //     flash.Store(&c.Controller)
    //     c.Redirect("/setting", 302)
    //     return
    // } else if b != nil {
    //     flash.Error("validation err!")
    //     flash.Store(&c.Controller)
    //     c.Redirect("/setting", 302)
    //     return
    // }
    // saveSetting(setting)
    // flash.Notice("Settings saved!")
    // flash.Store(&c.Controller)
    // c.Redirect("/setting", 302)
	
	
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	// users := models.GetAllUsers()
	// u.Data["json"] = users
	// u.ServeJSON()
	fmt.Println("111111111111")
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	fmt.Println("111111111111")
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	user, err := models.GetUser(uid)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = user
	// 	}
	// }
	// u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	fmt.Println("111111111111")
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	var user models.User
	// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// 	uu, err := models.UpdateUser(uid, &user)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = uu
	// 	}
	// }
	// u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	fmt.Println("111111111111")
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
	// u.Data["json"] = "delete success!"
	// u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	fmt.Println("111111111111")
	// username := u.GetString("username")
	// password := u.GetString("password")
	// if models.Login(username, password) {
	// 	u.Data["json"] = "login success"
	// } else {
	// 	u.Data["json"] = "user not exist"
	// }
	// u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}


