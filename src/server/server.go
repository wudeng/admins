package server

import (
	"github.com/labstack/echo"
	"data"
	"net/http"
	"net/url"
)



func Servers(c echo.Context) error {
	return c.JSON(http.StatusOK, data.H{"status": "ok", "data": data.ListServer()})
}

func CreateServer(c echo.Context) error {
	if c.FormValue("id") == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID不能为空"})
	}
	if c.FormValue("nm") == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器名字不能为空"})
	}
	resp, err := http.PostForm(data.SERVER_LIST + "add_server", url.Values{
		"server_id":{c.FormValue("id")},
		"name":{c.FormValue("nm")},
		"ip": {c.FormValue("ip")},
		"port": {c.FormValue("pt")},
		"http_port": {c.FormValue("http_pt")},
		"open_date": {"2018-01-10 10:00:00"},
	})
	defer resp.Body.Close()
	if err != nil {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器创建失败"})
	} else {
		adrd := &data.AdminRecord{
			AdminID: data.GetCurrentUserID(c),
			Kind:    data.OPERATE_KIND_ADD_SERVER,
			Target:  c.FormValue("id"),
		}
		adrd.Save()
		return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "服务器创建成功"})
	}
}

func Edit(c echo.Context) error {
	id := c.FormValue("id")
	if id == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID不能为空"})
	}
	name := c.FormValue("nm")
	if name == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID不能为空"})
	}
	resp, err := http.PostForm(data.SERVER_LIST + "edit_server", url.Values{
		"server_id":{id},
		"name":{name},
		"ip": {c.FormValue("ip")},
		"port": {c.FormValue("pt")},
		"http_port": {c.FormValue("http_pt")},
		"open_date": {"2018-01-10 10:00:00"},
	})
	defer resp.Body.Close()
	if err != nil {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器数据修改失败"})
	}
	return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "服务器更改成功"})
}


func Delete(c echo.Context) error {
	id := c.FormValue("id")
	if id != "" {
		resp, err := http.PostForm(data.SERVER_LIST + "delete_server", url.Values{
			"server_id":{id},
		})
		defer resp.Body.Close()
		if err != nil {
			return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "删除失败"})
		} else {
			return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "删除成功"})
		}
	} else {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器不能为空"})
	}
	return nil
}
