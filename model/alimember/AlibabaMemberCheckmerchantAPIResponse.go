package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberCheckmerchantAPIResponse 校验商家身份 API返回值
// alibaba.member.checkmerchant
//
// 校验商家身份
type AlibabaMemberCheckmerchantAPIResponse struct {
	model.CommonResponse
	AlibabaMemberCheckmerchantAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberCheckmerchantAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberCheckmerchantAPIResponseModel).Reset()
}

// AlibabaMemberCheckmerchantAPIResponseModel is 校验商家身份 成功返回结果
type AlibabaMemberCheckmerchantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_checkmerchant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberCheckmerchantAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnMessage = ""
	m.ReturnCode = ""
}

var poolAlibabaMemberCheckmerchantAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberCheckmerchantAPIResponse)
	},
}

// GetAlibabaMemberCheckmerchantAPIResponse 从 sync.Pool 获取 AlibabaMemberCheckmerchantAPIResponse
func GetAlibabaMemberCheckmerchantAPIResponse() *AlibabaMemberCheckmerchantAPIResponse {
	return poolAlibabaMemberCheckmerchantAPIResponse.Get().(*AlibabaMemberCheckmerchantAPIResponse)
}

// ReleaseAlibabaMemberCheckmerchantAPIResponse 将 AlibabaMemberCheckmerchantAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberCheckmerchantAPIResponse(v *AlibabaMemberCheckmerchantAPIResponse) {
	v.Reset()
	poolAlibabaMemberCheckmerchantAPIResponse.Put(v)
}
