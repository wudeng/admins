package main

import (
	"admincall"
	"basic/ssdb/gossdb"
	"data"

	"flag"
	"net/http"
	"operation"
	"role"
	"statistics"
	"user"

	_ "csv"

	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "errors"
	"server"
)
var (
	version = "v0.0.1"
)

// 登陆中间件
func loginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("login")
		if err != nil || cookie == nil || len(cookie.Value) <= 0 || data.Sessions.Get(cookie.Value) == nil {
			if c.Request().Method == "GET" {
				c.Request().Header.Add("Cache-Control", "no-cache")
				return c.Redirect(http.StatusTemporaryRedirect, "/login/login.html")
			} else if c.Request().Method == "POST" {
				return c.JSON(http.StatusOK, data.H{"status": "fail", "errorcode": 100, "msg": "未登陆"})
			}
			return err
		}
		return next(c)
	}
}

func main() {
	var config string
	flag.StringVar(&config, "conf", "./conf.json", "config path")
	flag.Parse()
	data.LoadConf(config)
	data.Version = version
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	//	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	e.Static("/assets", "AmazeUI/assets")

	e.Static("/users", "AmazeUI/users")
	e.Static("/operation", "AmazeUI/operation")
	e.Static("/roles", "AmazeUI/roles")
	e.Static("/servers", "AmazeUI/server")
	e.Static("/room", "AmazeUI/room")
	e.Static("/statistics", "AmazeUI/statistics")
	e.Static("/tools", "AmazeUI/tools")

	e.Static("/login", "AmazeUI/login")

	e.POST("/users/login", user.Login)
	e.POST("/users/logout", user.Logout)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/statistics/index.html?v="+data.Version)
	}, loginMiddleware)

	e.POST("/roles/list", role.List, loginMiddleware)
	e.POST("/roles/search", role.Search, loginMiddleware)
	e.POST("/roles/edit", role.Edit, loginMiddleware)


	e.POST("/users/create", user.Create, loginMiddleware)
	e.POST("/users/edit", user.Edit, loginMiddleware)
	e.POST("/users/list", user.List, loginMiddleware)
	e.POST("/users/delete", user.Delete, loginMiddleware)
	e.POST("/users/getdetail", user.GetSelfDetail, loginMiddleware)
	e.POST("/users/record", user.Record, loginMiddleware)

	e.POST("/users/getloginlimit", user.GetLoginLimit, loginMiddleware)
	e.POST("/users/delloginlimit", user.DelLoginLimit, loginMiddleware)

	e.POST("/servers/list", server.Servers, loginMiddleware)
	e.POST("/servers/create", server.CreateServer, loginMiddleware)
	e.POST("/servers/edit", server.Edit, loginMiddleware)
	e.POST("/servers/delete", server.Delete, loginMiddleware)

	e.POST("/group/create", user.CreateGroup, loginMiddleware)
	e.POST("/group/edit", user.EditGroup, loginMiddleware)
	e.POST("/group/list", user.Groups, loginMiddleware)
	e.POST("/group/delete", user.DeleteGroup, loginMiddleware)

	e.POST("/roles/listonline", role.ListOnline, loginMiddleware) //  在线玩家列表

	e.POST("/operation/privaterecord", operation.PrivateRecord, loginMiddleware)       // 私人局牌局记录
	e.POST("/operation/matchrecord", operation.MatchRecord, loginMiddleware)           // 比赛场牌局记录
	e.POST("/operation/normalrecord", operation.NormalRecord, loginMiddleware)         //  金币场牌局记录
	e.POST("/operation/cardrecord", operation.CardRecode, loginMiddleware)             //  金币场牌打牌记录
	e.POST("/operation/issueprops", operation.IssueProps, loginMiddleware)             //
	e.POST("/operation/postbox", operation.Postbox, loginMiddleware)                   //
	e.POST("/operation/issuelist", operation.IssuePropsList, loginMiddleware)          //
	e.POST("/operation/loginrecord", operation.LoginRecord, loginMiddleware)           //
	e.POST("/operation/roomcreaterecord", operation.RoomCreateRecord, loginMiddleware) // 私人房创建记录
	e.POST("/operation/consume", operation.Consume, loginMiddleware)

	e.POST("/operation/charge", operation.GetChargeOrder, loginMiddleware)    // 下单记录
	e.POST("/operation/transition", operation.GetTransition, loginMiddleware) // 交易记录

	e.POST("/statistics/online", statistics.Online, loginMiddleware)
	e.POST("/statistics/newuser", statistics.NewUser, loginMiddleware)
	e.POST("/statistics/index", statistics.Index, loginMiddleware)

	e.POST("/tools/getnotice", operation.GetNotice, loginMiddleware)
	e.POST("/tools/issuenotice", operation.AddNotice, loginMiddleware)
	conndb()


	data.InitAdmin()
	data.InitServers()
	data.LoadLimitIPs()
	admincall.Init(data.Conf.CallServer)
	e.Start(data.Conf.Port)
}

// 链接数据库
func conndb() {
	glog.Infoln("Config: ", data.Conf)
	gossdb.Connect(data.Conf.Db.Ip, data.Conf.Db.Port, data.Conf.Db.Thread)
	defer glog.Flush()
}
