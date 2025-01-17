package controller

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/kataras/iris"
	"hdt_app_go/common"
	"hdt_app_go/common/hdtcodec"
	. "hdt_app_go/gateway_app/log"
	. "hdt_app_go/gateway_app/rpc"
	proto "hdt_app_go/protcol"
	"io/ioutil"
)

var ServerTag = "errcode"

func ECHO(ctx iris.Context) {
	ctx.JSON("hello")
}

type RegisterReq struct {
	Tel          string `json:"tel"`           //注册手机号
	Pwd          string `json:"pwd"`           //密码
	Code         string `json:"code"`          //验证码
	RegisterFrom int32  `json:"register_from"` //注册来源：1-安卓，2-IOS
}

type LoginReq struct {
	Tel string `json:"tel"` //注册手机号
	Pwd string `json:"pwd"` //密码
}

type FindPwdReq struct {
	Tel  string `json:"tel"` //注册手机号
	Pwd  string `json:"pwd"` //密码
	Code string `json:"code"`
}

func Register(ctx iris.Context) {
	data := map[string]interface{}{ //定义返回的map
		"errcode": proto.ERR_OK,
	}

	ip := ctx.RemoteAddr() //获取请求段IP地址

	body, _ := ioutil.ReadAll(ctx.Request().Body)    //读取body
	jsonStr, _ := hdtcodec.HdtDecodeV0(string(body)) //根据编码规则解析请求的body
	//jsonStr, _ := ioutil.ReadAll(ctx.Request().Body)

	var req RegisterReq
	err := json.Unmarshal([]byte(jsonStr), &req) //根据请求参数结构体解析参数
	if err != nil || len(req.Tel) < 11 || req.RegisterFrom <= 0 ||
		req.RegisterFrom > 2 { //手机号不能小于11位，注册来源只能（1.安卓，2.IOS）
		Log.Err(err.Error())
		data["errcode"] = proto.ERR_PARAM
		data["err_msg"] = err.Error()
		ctx.JSON(data)
		return
	}

	result := RpcClient.Register.QianXunSnsVerify(req.Tel, req.Code) //验证短信验证码
	if result == proto.ERR_OK {                                      //如果验证成功[匹配成功]则创建账号
		data["errcode"] = RpcClient.Register.CreateAccountByTel(req.Tel, req.Pwd, ip, req.RegisterFrom)
	} else {
		data["errcode"] = result
	}

	ctx.JSON(data)
	return
}

func FindPwdByTel(ctx iris.Context) {
	data := map[string]interface{}{
		"errcode": proto.ERR_OK,
	}

	//以下这段代码是对传入的参数进行校验
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	jsonStr, _ := hdtcodec.HdtDecodeV0(string(body))
	//jsonStr, _ := ioutil.ReadAll(ctx.Request().Body)
	var req FindPwdReq
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		Log.Err(err.Error())
		data["errcode"] = proto.ERR_PARAM
		data["err_msg"] = err.Error()
		ctx.JSON(data)
		return
	}
	valid := validation.Validation{}
	valid.Required(req.Tel, "tel")
	valid.Required(req.Pwd, "pwd")
	valid.Required(req.Code, "code")
	if valid.HasErrors() {
		for _, err1 := range valid.Errors {
			Log.Warningf("invalid args %s: %s", err1.Key, err1.Message)
		}
		data["errcode"] = proto.ERR_PARAM
		ctx.JSON(data)
		return
	}

	result := RpcClient.Register.QianXunSnsVerify(req.Tel, req.Code)
	if result == proto.ERR_OK {
		data["errcode"] = RpcClient.Register.ModifyPwdByTel(req.Tel, req.Pwd)
	} else {
		data["errcode"] = result
	}

	ctx.JSON(data)
	return
}

func Login(ctx iris.Context) {
	data := map[string]interface{}{
		"errcode": proto.ERR_OK,
	}

	//以下这段代码是对传入的参数进行校验
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	jsonStr, _ := hdtcodec.HdtDecodeV0(string(body))
	//jsonStr, _ := ioutil.ReadAll(ctx.Request().Body)
	var req LoginReq
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		Log.Err(err.Error())
		data["errcode"] = proto.ERR_PARAM
		data["err_msg"] = err.Error()
		ctx.JSON(data)
		return
	}
	valid := validation.Validation{}
	valid.Required(req.Tel, "tel")
	valid.Required(req.Pwd, "pwd")
	if valid.HasErrors() {
		for _, err1 := range valid.Errors {
			Log.Warningf("invalid args %s: %s", err1.Key, err1.Message)
		}
		data["errcode"] = proto.ERR_PARAM
		ctx.JSON(data)
		return
	}

	token := common.GenUserToken(req.Tel)           //生成token
	RpcClient.Register.SetUserToken(req.Tel, token) //设置token

	errCode, userInfo := RpcClient.Register.Login(req.Tel, req.Pwd) //获取用户信息
	data["errcode"] = errCode
	if errCode == proto.ERR_OK {
		data["user_info"] = userInfo
		data["token"] = token
	}

	ctx.JSON(data)
	return
}
