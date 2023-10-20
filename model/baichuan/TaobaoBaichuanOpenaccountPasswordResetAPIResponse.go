package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountpasswordresetAPIResponse 百川找回密码 API返回值
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
type TaobaobaichuanopenaccountpasswordresetAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanopenaccountpasswordresetAPIResponseModel
}

// TaobaobaichuanopenaccountpasswordresetAPIResponseModel is 百川找回密码 成功返回结果
type TaobaobaichuanopenaccountpasswordresetAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_password_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
