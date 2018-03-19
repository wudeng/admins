(function ($) {
    'use strict';
    $(function () {
        var $fullText = $('.admin-fullText');
        $('#admin-fullscreen').on('click', function () {
            $.AMUI.fullscreen.toggle();
        });

        $(document).on($.AMUI.fullscreen.raw.fullscreenchange, function () {
            $fullText.text($.AMUI.fullscreen.isFullscreen ? '退出全屏' : '开启全屏');
        });
    });
})(jQuery);
//去掉字符串前后所有空格
function Trim(str)
{
    return str.replace(/(^\s*)|(\s*$)/g, "");
}


function geturl(url) {
    return url+"?v="+getcookie("version")
}

//获取指定名称的cookie的值
function getcookie(objname) {
    var arrstr = document.cookie.split("; ");
    for (var i = 0; i < arrstr.length; i++) {
        var temp = arrstr[i].split("=");
        if (temp[0] == objname) return unescape(temp[1]);
    }
}
function onlogout() {
    $.Confirm("你要退出系统吗？",function (bool) {
        if (bool){
            $.post({

                type: "POST",
                url: "/users/logout",
                data: {},
                dataType: "json",
                success: function (data) {
                   location.href ="/login/login.html"
                }
            });
        }
    })
}
function add0(m){return m<10?'0'+m:m }
function format(shijianchuo)
{
//shijianchuo是整数，否则要parseInt转换
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y+'-'+add0(m)+'-'+add0(d)+' '+add0(h)+':'+add0(mm)+':'+add0(s);
}

function formatData(shijianchuo)
{
//shijianchuo是整数，否则要parseInt转换
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();

    return y+'-'+add0(m)+'-'+add0(d);
}

function ipToNumber(ip) {
    var num = 0;
    if(ip == "") {
        return num;
    }
    var aNum = ip.split(".");
    if(aNum.length != 4) {
        return num;
    }
    num += parseInt(aNum[0]) << 24;
    num += parseInt(aNum[1]) << 16;
    num += parseInt(aNum[2]) << 8;
    num += parseInt(aNum[3]) << 0;
    num = num >>> 0;//这个很关键，不然可能会出现负数的情况
    return num;
}

function numberToIp(number) {
    var ip = "";
    if(number <= 0) {
        return ip;
    }
    var ip3 = (number << 0 ) >>> 24;
    var ip2 = (number << 8 ) >>> 24;
    var ip1 = (number << 16) >>> 24;
    var ip0 = (number << 24) >>> 24

    ip += ip3 + "." + ip2 + "." + ip1 + "." + ip0;

    return ip;
}

//调用方法 如
//post('pages/statisticsJsp/excel.action', {html :prnhtml,cm1:'sdsddsd',cm2:'haha'});
function post(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none";
    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        // alert(opt.name)
        temp.appendChild(opt);
    }
    document.body.appendChild(temp);
    temp.submit();
    return temp;
}

function getHeadBar() {
    return [
        {
            name:  "管理员",
            icon: "am-icon-users",
            items:    [
                {icon:"am-icon-group",path:geturl("/users/group_list.html"), name:"用户组管理"},
                {icon:"am-icon-user",path:geturl("/users/list.html"), name:"后台用户账号管理"},
                {icon:"am-icon-eye",path:geturl("/users/record.html"), name:"后台操作日志"},
                {icon:"am-icon-edit",path:geturl("/users/edit.html"), name:"修改我的密码"}]
        },
        {
            name:  "系统设置",
            icon: "am-icon-bars",
            items:    [
         /*       {icon:"am-icon-user",path:"#", name:"资料"},*/
                {icon:"am-icon-user",path:geturl("/users/loginlimit.html"), name:"恶意登陆IP"},
                {icon:"am-icon-sign-out",path:"#",action:"javascript:onlogout()", name:"退出"}]
        }
    ]
}


function getSideBar() {
    return [
        {
            name:"服务器管理",
            icon:"am-icon-user",
            items: [
                {icon:"am-icon-table", path:geturl("/servers/list.html"), name:"服务器列表"}
            ]
        },
        {
            name:"玩家管理",
            icon:"am-icon-user",
            items: [
            { icon:"am-icon-table",path:geturl("/roles/list.html"),  name:"玩家列表"}

          /*
           { icon:"am-icon-table",path:geturl("/roles/feedback.html"),  name:"反馈"}
          {icon:"am-icon-table",path:geturl("/roles/listonline.html"),name: "在线玩家"},
            {icon:"am-icon-table",path:geturl("/roles/gainrank.html"), name:"每日盈利排名"},
            {icon:"am-icon-table",path:geturl("/roles/winrank.html"), name:"胜局排名"},
            {icon:"am-icon-table",path:geturl("/roles/coinrank.html"), name:"等级排名"},
            {icon:"am-icon-table",path:geturl("/roles/levelrank.html"),name: "等级排名"}*/
            ]
        },
        {
            name:  "邮件和奖励",
            icon:"am-icon-drupal",
            items:    [
          /*  {icon:"am-icon-table",path:geturl("/operation/provide.html"), name:"道具/钻石发放"},
            {icon:"am-icon-table",path:geturl("/operation/providerecord.html"), name:"发放记录"},*/
            {icon:"am-icon-table",path:geturl("/operation/postbox.html"), name:"发送邮件和奖励"}
 /*     {icon:"am-icon-table",path:geturl("/operation/emaillist.html", name:"邮件记录"}*/
            ]
        },
        {
            name: "运营工具",
            icon:"am-icon-eye",
            items:      [
                {icon:"am-icon-table",path:geturl("/tools/notice.html"), name:"游戏公告"}
              ]
        },
        /*{
            name: "日志管理",
            icon:"am-icon-eye",
            items:      [
                {icon:"am-icon-table",path:geturl("/operation/privaterecord.html"), name:"私人局记录"},
                {icon:"am-icon-table",path:geturl("/operation/matchrecord.html"),name: "比赛场记录"},
                {icon:"am-icon-table",path:geturl("/operation/normalrecord.html"), name:"金币场记录"},
                {icon:"am-icon-table",path:geturl("/operation/privatecreate.html"), name:"私人房创建记录"},
                { icon:"am-icon-table",path:geturl("/operation/consume.html"),  name:"消耗日志"}
                //{icon:"am-icon-table",path:geturl("/operation/exchangerecord.html"), name:"虚拟兑换记录"},
                //{icon:"am-icon-table",path:geturl("/operation/exchangerecord.html"), name:"实物兑换记录"},
                //{icon:"am-icon-table",path:geturl("/operation/loginrecord.html"),name: "登录日志"},
            ]
        },
        {
            name: "订单管理",
            icon:"am-icon-shield",
            items:     [
            {icon:"am-icon-table",path:geturl("/operation/chargeorder.html"), name:"下单记录"},
            {icon:"am-icon-table",path:geturl("/operation/transition.html"), name:"交易记录"}
            ]
        },*/
        {
            name: "数据统计",
            icon:"am-icon-file",
            items:     [
                   {icon:"am-icon-table",path:geturl("/statistics/online.html"),name: "在线"}
               /* {icon:"am-icon-table",path:geturl("/statistics/newuser.html"),name: "新增"}*/
                /*  {icon:"am-icon-table",path:geturl("/statistics/active.html"),name: "活跃"},
            {icon:"am-icon-table",path:geturl("/statistics/remainder.html"),name: "留存"}*/
                ]
        }
    ]
}

(function(){
    Vue.component('sidebar-component', {
        template:`
            <!-- sidebar start -->
            <div class="tpl-left-nav tpl-left-nav-hover">
                <div class="tpl-left-nav-list">
                    <ul class="tpl-left-nav-menu" id="admin-sidebar-list">
                        <li class="tpl-left-nav-item">
                            <a href="/" class="nav-link active">
                                <i class="am-icon-home"></i>
                                <span>首页</span>
                            </a>
                        </li>
                        <!-- table start -->

                        <!-- table start -->

                        <li v-for="item in items" class="tpl-left-nav-item">
                            <a href="javascript:;" class="nav-link tpl-left-nav-link-list">
                                <i :class="item.icon"></i>
                                <span>{{ item.name }} </span>
                                <i class="am-icon-angle-right tpl-left-nav-more-ico am-fr am-margin-right"></i>
                            </a>
                            <ul class="tpl-left-nav-sub-menu" style="display: block;">
                                <li v-for="i in item.items">
                                    <a :href="i.path">
                                        <i :class="i.icon"></i>
                                        <span>{{ i.name}}</span>
                                        <i class="am-icon-star tpl-left-nav-content-ico am-fr am-margin-right"></i>
                                    </a>
                                </li>
                            </ul>
                        </li>

                    </ul>
                </div>
            </div>
        `,
        data: function() {
            return {
                items: getSideBar()
            }
        }
    })
})();

(function(){
    Vue.component('header-component', {
        template: `
            <header class="am-topbar am-topbar-inverse admin-header">
                <div class="am-topbar-brand">
                    <strong>天书</strong>
                    <small>后台管理系统</small>
                </div>
                <div class="am-icon-list tpl-header-nav-hover-ico am-fl am-margin-right"></div>

                <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only"
                        data-am-collapse="{target: '#topbar-collapse'}"><span class="am-sr-only">导航切换</span> <span
                        class="am-icon-bars"></span></button>
                <!-- headbar start -->
                <div class="am-collapse am-topbar-collapse" id="topbar-collapse">
                    <ul class="am-nav am-nav-pills am-topbar-nav am-topbar-right admin-header-list tpl-header-list"
                        id="admin-head-list">
                        <li v-for="item in items" class="am-dropdown" data-am-dropdown data-am-dropdown-toggle>
                            <a class="am-dropdown-toggle tpl-header-list-link" href="javascript:;">
                                <span :class="item.icon"></span> {{item.name}} <span class="am-icon-caret-down"></span>
                            </a>
                            <ul class="am-dropdown-content tpl-dropdown-content">
                                <li v-for="i in item.items">
                                    <a :href="i.path" type="button" :onclick="i.action"><span :class="i.icon"></span> </span> {{
                                        i.name}}</a></li>

                            </ul>
                        </li>
                    </ul>
                </div>
                <!-- headbar end -->
            </header>
        `,
        data: function() {
            return {
                items: getHeadBar()
            }
        }
    })
})();

function getServerList(self) {
    $.ajax({
        type: "POST",
        url: "/servers/list",
        data: {},
        dataType: "json",
        success: function (data) {
            if (data["status"] == "ok") {
                //console.info(data["data"])
                self.servers = data.data
            }
        }
    })
}

(function() {
    Vue.component('server-list-component', {
        props: ["servers", "content"],
        template: `
              <div  class="am-form-group">
                  <label  class="am-u-sm-3 am-form-label">选择服务器</label>
                  <div class="am-u-sm-9">
                    <div v-for="item in servers">
                        <div class="am-u-sm-2">
                              <input type="checkbox" v-bind:value="item.id" v-model="content.server_ids"></input>
                              <label >{{item.nm}}</label>
                        </div>
                    </div>
                </div>
              </div>
        `
    })
})();