package damaiticklet

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaitickletqrcodedecodeAPIResponse 票夹-动态二维码-解码 API返回值
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
type AlibabadamaitickletqrcodedecodeAPIResponse struct {
	model.CommonResponse
	AlibabadamaitickletqrcodedecodeAPIResponseModel
}

// AlibabadamaitickletqrcodedecodeAPIResponseModel is 票夹-动态二维码-解码 成功返回结果
type AlibabadamaitickletqrcodedecodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_ticklet_qrcode_decode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
