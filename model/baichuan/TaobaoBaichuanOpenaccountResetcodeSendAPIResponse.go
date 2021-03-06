package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountResetcodeSendAPIResponse 百川发送找回密码验证码 API返回值
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel
}

// TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel is 百川发送找回密码验证码 成功返回结果
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_resetcode_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
