package damaiticklet

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiTickletQrcodeDecodeAPIResponse 票夹-动态二维码-解码 API返回值
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
type AlibabaDamaiTickletQrcodeDecodeAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiTickletQrcodeDecodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiTickletQrcodeDecodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiTickletQrcodeDecodeAPIResponseModel).Reset()
}

// AlibabaDamaiTickletQrcodeDecodeAPIResponseModel is 票夹-动态二维码-解码 成功返回结果
type AlibabaDamaiTickletQrcodeDecodeAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaDamaiTickletQrcodeDecodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = ""
	m.IsSuccess = false
}

var poolAlibabaDamaiTickletQrcodeDecodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiTickletQrcodeDecodeAPIResponse)
	},
}

// GetAlibabaDamaiTickletQrcodeDecodeAPIResponse 从 sync.Pool 获取 AlibabaDamaiTickletQrcodeDecodeAPIResponse
func GetAlibabaDamaiTickletQrcodeDecodeAPIResponse() *AlibabaDamaiTickletQrcodeDecodeAPIResponse {
	return poolAlibabaDamaiTickletQrcodeDecodeAPIResponse.Get().(*AlibabaDamaiTickletQrcodeDecodeAPIResponse)
}

// ReleaseAlibabaDamaiTickletQrcodeDecodeAPIResponse 将 AlibabaDamaiTickletQrcodeDecodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiTickletQrcodeDecodeAPIResponse(v *AlibabaDamaiTickletQrcodeDecodeAPIResponse) {
	v.Reset()
	poolAlibabaDamaiTickletQrcodeDecodeAPIResponse.Put(v)
}
