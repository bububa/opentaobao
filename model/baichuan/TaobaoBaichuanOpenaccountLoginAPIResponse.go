package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountloginAPIResponse 百川用户名密码登录 API返回值
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
type TaobaobaichuanopenaccountloginAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanopenaccountloginAPIResponseModel
}

// TaobaobaichuanopenaccountloginAPIResponseModel is 百川用户名密码登录 成功返回结果
type TaobaobaichuanopenaccountloginAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
