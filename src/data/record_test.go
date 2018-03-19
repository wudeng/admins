package data

import (
	"testing"
	"net/http"
	"net/url"
	"fmt"
	"net/http/httptest"
	"strings"
	"github.com/labstack/echo"
	"strconv"
)

func Test_get_server(t *testing.T) {
	var rawParams string = "sids%5B%5D=2&sids%5B%5D=1&Id=1&WidgetType=1&Count=1&Desc=&UserIds=&VIP=22&Kind=1"
	values, _ := url.ParseQuery(rawParams)
	fmt.Println(values["sids[]"])
	fmt.Println(values["Id"])
	InitServers()
	for _, sid := range values["sids[]"] {
		fmt.Println(sid)
		var s = Server{}
		s.Id, _ = strconv.ParseInt(sid, 10, 64)
		s.Get()
		fmt.Println(s)
	}

}
func Test_normal_record(t *testing.T) {
	//gossdb.Connect("192.168.50.10", 8888, 1)
	////InitServers()
	//var s Server
	//s.Id = 2
	//s.Get()
	//ret := s.Mail("1")
	//fmt.Println(ret)
	resp, err := http.PostForm(SERVER_LIST+"add_server", url.Values{
		"server_id":{"1000"},
		"name":{"测试2服"},
		"ip": {"127.0.0.1"},
		"port": {"8899"},
		"http_port": {"18899"},
		"open_date": {"2018-01-10 10:00:00"},
	})
	if err != nil {

	}
	defer resp.Body.Close()



	//fmt.Println(ListServer())
}
func Test_delete_server(t *testing.T) {
	resp, err := http.PostForm(SERVER_LIST+"delete_server", url.Values{
		"server_id":{"1000"},
	})
	if err != nil {

	}
	defer resp.Body.Close()
}

func Test_encoding(t *testing.T) {
	v := url.Values{}
	v.Set("server_id", "1000")
	v.Add("name", "测")
	fmt.Println(v.Get("server_id"))
	fmt.Println(v.Get("name"))
	fmt.Println(v.Encode())
}

func Test_params(t *testing.T) {
	f := make(url.Values)
	f.Set("name", "Jon Snow")
	f.Add("name", "Nerd Stark")
	f.Set("email", "jon@labstack.com")

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
	c := e.NewContext(req, nil)

	fmt.Println(f["name"])
	fmt.Println(c.FormValue("name"))
	fmt.Println(c.FormValue("email"))
	// FormValue


	//// FormParams
	//params, err := c.FormParams()
	//if assert.NoError(t, err) {
	//	assert.Equal(t, url.Values{
	//		"name":  []string{"Jon Snow"},
	//		"email": []string{"jon@labstack.com"},
	//	}, params)
	//}
}
