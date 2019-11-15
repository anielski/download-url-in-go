package app

import (
	"fmt"
	"github.com/anielski/download-url-in-go/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/random"
	"gl.tzv.io/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

// API godoc
var API *Application

// Application godoc
type Application struct {
	Server  *echo.Echo  `json:"server"`
	Name    string      `json:"name"`
	Version string      `json:"version"`
	ENV     string      `json:"env"`
	Config  viper.Viper `json:"application_config"`
	DB      *gorm.DB     `json:"db"`
}

// CustomValidator godoc
type CustomValidator struct {
	validator *validator.Validate
}

// Validate godoc
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Logger godoc
func Logger(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05 -0700 MST"),
		})
	}

	req := c.Request()
	res := c.Response()

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	return logrus.WithFields(logrus.Fields{
		"id":     id,
		"at":     time.Now().Format("2006-01-02 15:04:05 -0700 MST"),
		"method": req.Method,
		"uri":    req.URL.String(),
		"ip":     req.RemoteAddr,
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Logger(c).Info("incoming request")
		return next(c)
	}
}

func connDB() {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",  os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	//Logger(nil).Infof("CONNECT %s", dbSource)
	db := models.NewDB(dbSource)
	API.DB = db
}

func bindDb() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", API.DB)
			return next(c)
		}
	}
}


// Init godoc
func Init() *echo.Echo {
	logrus.SetLevel(logrus.DebugLevel)
	// Echo instance
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Error Handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
			report.Code = http.StatusBadRequest
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return random.String(32)
		},
	}))

	//MAX 1MB payload
	e.Use(middleware.BodyLimit("1M"))

	// Middleware
	e.Use(middlewareLogging)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${id}] method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	//DB
	connDB()
	e.Use(bindDb())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Assign to global
	return e
}
