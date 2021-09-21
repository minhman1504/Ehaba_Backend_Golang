package handler

import (
	"ehaba_backend_golang/db"
	"ehaba_backend_golang/model"
	"ehaba_backend_golang/model/req"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	MongoClient db.MongoDB
}

func (u *UserHandler) SignIn(c echo.Context) error {
	req := req.ReqSignIn{}               //	tạo đối tượng để hứng giá trị được gửi lên
	if err := c.Bind(&req); err != nil { //run when err != nil 				chạy khi có lỗi xảy ra
		fmt.Println("Error-SignIn-Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} //Đến đây thì có thể lấy được các giá trị của req (chú ý struct phải có tag để khi nhận được json có thể convert lại thành struct)

	//_, err := u.UserRepo.CheckLogin(c.Request().Context(), req)

	//fmt.Println(req.Email, req.Password, err)

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Request Success",
		Data:       req.Email + req.Password,
	})
}

func (u *UserHandler) SignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil { //run when err != nil 				chạy khi có lỗi xảy ra
		fmt.Println("Error-SignUp-Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} //Đến đây thì có thể lấy được các giá trị của req (chú ý struct phải có tag để khi nhận được json có thể convert lại thành struct)

	//check email

	if u.MongoClient.Email_IsOnDatabase(c.Request().Context(), req.Email) {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Email đã đăng ký",
			Data:       nil,
		})
	}

	//create user to store into database
	//create uuid
	userId, _ := uuid.NewUUID()
	user := model.User{
		UserId:    userId.String(),
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
		Birthday:  req.Birthday,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
	}

	if !u.MongoClient.InsertOneUser(c.Request().Context(), user) {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Lưu vào cơ sở dữ liệu bị lỗi",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng ký tài khoản thành công",
		Data:       req.Email,
	})
}
