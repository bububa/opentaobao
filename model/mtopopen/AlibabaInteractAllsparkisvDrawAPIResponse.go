package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractAllsparkisvDrawAPIResponse allspark提供抽奖tida接口对应鉴权接口 API返回值
// alibaba.interact.allsparkisv.draw
//
// 该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
type AlibabaInteractAllsparkisvDrawAPIResponse struct {
	model.CommonResponse
	AlibabaInteractAllsparkisvDrawAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractAllsparkisvDrawAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractAllsparkisvDrawAPIResponseModel).Reset()
}

// AlibabaInteractAllsparkisvDrawAPIResponseModel is allspark提供抽奖tida接口对应鉴权接口 成功返回结果
type AlibabaInteractAllsparkisvDrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_allsparkisv_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ddd
	Ddd string `json:"ddd,omitempty" xml:"ddd,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractAllsparkisvDrawAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ddd = ""
}

var poolAlibabaInteractAllsparkisvDrawAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractAllsparkisvDrawAPIResponse)
	},
}

// GetAlibabaInteractAllsparkisvDrawAPIResponse 从 sync.Pool 获取 AlibabaInteractAllsparkisvDrawAPIResponse
func GetAlibabaInteractAllsparkisvDrawAPIResponse() *AlibabaInteractAllsparkisvDrawAPIResponse {
	return poolAlibabaInteractAllsparkisvDrawAPIResponse.Get().(*AlibabaInteractAllsparkisvDrawAPIResponse)
}

// ReleaseAlibabaInteractAllsparkisvDrawAPIResponse 将 AlibabaInteractAllsparkisvDrawAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractAllsparkisvDrawAPIResponse(v *AlibabaInteractAllsparkisvDrawAPIResponse) {
	v.Reset()
	poolAlibabaInteractAllsparkisvDrawAPIResponse.Put(v)
}
