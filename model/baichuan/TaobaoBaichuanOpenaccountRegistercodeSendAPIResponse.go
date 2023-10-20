package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse 百川发送注册验证码 API返回值
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
type TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel
}

// TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel is 百川发送注册验证码 成功返回结果
type TaobaoBaichuanOpenaccountRegistercodeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_registercode_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
