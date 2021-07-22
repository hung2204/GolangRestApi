package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

//muon sd phai dang ki model voi orm
func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id    int    `orm:"auto" json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

var listUsers = []User{
	{
		Name: "Hung",
		Age:  23,
	},
	{
		Name: "Huan",
		Age:  23,
	},
	{
		Name: "Linh",
		Age:  23,
	},
	{
		Name: "Kien",
		Age:  23,
	},
}

func GetAllUsers(c echo.Context) error {
	//ghi header tra ve content dang JSON
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//bao cho Response ghi StatusOK => client nhan status 200
	c.Response().WriteHeader(http.StatusOK)

	//tao JSON enc de ghi vao Response
	enc := json.NewEncoder(c.Response())

	// dung enc ghi vao trong Response
	for _, user := range listUsers {
		if err := enc.Encode(user); err != nil {
			return err
		}
		//Response().Flush(): gui data ve client
		c.Response().Flush()
		time.Sleep(2 * time.Second)
	}
	return nil
}

func CreateUser(c echo.Context) error {
	//goi user
	user := &User{}
	//ktra loi
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind user error: %v", err)
		return err
	}
	//tao moi NewOrm
	o := orm.NewOrm()
	id, err := o.Insert(user)
	//ktra loi
	if err != nil {
		glog.Errorf("insert user error: %v", err)
		return err
	}
	//luu vao database
	glog.Infof("insert at row: %v", id)
	return c.JSON(http.StatusOK, user)
}

func ReadUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	// name := c.QueryParam("name")
	o := orm.NewOrm()
	user := &User{
		Id: id,
		// Name: name,
	}
	err := o.Read(user, "id")
	if err != nil {
		glog.Errorf("get user %s error: %v", id, err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind user error: %v", err)
		return err
	}
	glog.Infof("req update user: %+v", user)
	o := orm.NewOrm()
	_, err := o.Update(user, "Name", "Age", "Phone")
	if err != nil {
		glog.Errorf("update user %s error: %v", user, err)
		return err
	}
	userUpdate := &User{
		Name: user.Name,
	}
	o.Read(userUpdate, "Name")
	return c.JSON(http.StatusOK, userUpdate)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	glog.Infof("Deleting user %d", id)
	user := &User{
		Id: id,
		// Name: name,
	}
	o := orm.NewOrm()
	row, err := o.Delete(user)
	if err != nil {
		glog.Errorf("delete user %d error %v\n", id, err)
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("delete user id: %d at %d", id, row))
}
