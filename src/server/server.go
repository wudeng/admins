package server

import (
	"github.com/labstack/echo"
	"data"
	"net/http"
	"strconv"
	"github.com/golang/glog"
)

func Servers(c echo.Context) error {
	list := data.ListServer()
	return c.JSON(http.StatusOK, data.H{"status": "ok", "data": list})
}

func CreateServer(c echo.Context) error {
	if c.FormValue("id") == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID不能为空"})
	}
	if c.FormValue("nm") == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器名字不能为空"})
	}
	s := &data.Server{}
	s.Id, _ = strconv.ParseInt(c.FormValue("id"), 10, 64)
	s.Nm = c.FormValue("nm")
	s.Ip = c.FormValue("ip")
	Pt, _ := strconv.ParseInt(c.FormValue("pt"), 10, 64)
	Re, _ := strconv.ParseInt(c.FormValue("re"), 10, 64)
	s.Pt = uint16(Pt)
	s.Re = uint8(Re)
	s.St = 0

	glog.Infoln("Id:", s.Id, "Nm:", s.Nm, "Ip:", s.Ip, "Port:", s.Pt)

	if s.Exist() {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID已经存在"})
	} else {
		adrd := &data.AdminRecord{
			AdminID: data.GetCurrentUserID(c),
			Kind:    data.OPERATE_KIND_ADD_SERVER,
			Target:  strconv.FormatInt(s.Id, 10),
		}
		adrd.Save()

		err := s.Save()
		if err != nil {
			return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器创建失败"})
		} else {

			return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "服务器创建成功"})
		}

	}
	return nil
}

func Edit(c echo.Context) error {
	if c.FormValue("id") == "" {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器ID不能为空"})
	}
	s := &data.Server{}
	s.Id, _ = strconv.ParseInt(c.FormValue("id"), 10, 64)

	err := s.Get()
	if err != nil {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器不存在"})
	}

	if c.FormValue("nm") != "" {
		adrd := &data.AdminRecord{
			AdminID: data.GetCurrentUserID(c),
			Kind:    data.OPERATE_KIND_MODIFY_SERVER,
			Target:  strconv.FormatInt(s.Id, 10),
			Desc:    "名字",
			After:   c.FormValue("nm"),
			Pre:     s.Nm,
		}
		adrd.Save()

		s.Nm = c.FormValue("nm")
	}
	if c.FormValue("ip") != "" {
		adrd := &data.AdminRecord{
			AdminID: data.GetCurrentUserID(c),
			Kind:    data.OPERATE_KIND_MODIFY_ADMIN,
			Target:  strconv.FormatInt(s.Id, 10),
			Pre:     s.Ip,
			After:   c.FormValue("ip"),
			Desc:    "IP",
		}
		adrd.Save()

		s.Ip = c.FormValue("ip")
	}
	if c.FormValue("pt") != "" {
		adrd := &data.AdminRecord{
			AdminID: data.GetCurrentUserID(c),
			Kind:    data.OPERATE_KIND_MODIFY_ADMIN,
			Target:  strconv.FormatInt(s.Id, 10),
			Pre:     strconv.FormatInt(int64(s.Pt), 10),
			After:   c.FormValue("pt"),
			Desc:    "所属组",
		}
		adrd.Save()

		Pt, _ := strconv.ParseInt(c.FormValue("pt"), 10, 64)
		s.Pt = uint16(Pt)
	}

	if err := s.Save(); err != nil {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器数据修改失败"})
	}
	return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "服务器更改成功"})
}


func Delete(c echo.Context) error {
	s := &data.Server{}
	if c.FormValue("id") != "" {
		s.Id, _ = strconv.ParseInt(c.FormValue("id"), 10, 64)
		if s.Del() == nil {
			adrd := &data.AdminRecord{
				AdminID: data.GetCurrentUserID(c),
				Kind:    data.OPERATE_KIND_DEL_SERVER,
				Target:  strconv.FormatInt(s.Id, 10),
			}
			adrd.Save()

			return c.JSON(http.StatusOK, data.H{"status": "ok", "msg": "删除成功"})
		} else {
			return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "删除失败"})
		}
	} else {
		return c.JSON(http.StatusOK, data.H{"status": "fail", "msg": "服务器不能为空"})
	}
	return nil
}
