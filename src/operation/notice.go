package operation

import (
	"data"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"fmt"
	"github.com/golang/glog"
)

func GetNotice(c echo.Context) error {
	list := data.GetNotice()
	if len(list) > 0 {
		return c.JSON(http.StatusOK, data.H{"status": "ok", "data": list[0]})
	}
	return c.JSON(http.StatusOK, data.H{"status": "ok", "data": &data.Notice{}})
}
func AddNotice(c echo.Context) error {
	values, _ := c.FormParams()
	serverIds := values["server_ids[]"]
	content := values.Get("Content")
	s := data.Server{}
	for _, serverId := range serverIds {
		s.Id, _ = strconv.ParseInt(serverId, 10, 64)
		s.Get()
		fmt.Println(s)
		go s.Notice(content)
	}
	//kind, _ := strconv.Atoi(c.FormValue("Kind"))                 // uint32
	//expire, _ := strconv.ParseInt(c.FormValue("Expire"), 10, 64) // uint32
	glog.Infoln(serverIds, content)
	//data.AddNotice(uint32(kind), expire, title, content)

	return c.JSON(http.StatusOK, data.H{"status": "ok"})
}
