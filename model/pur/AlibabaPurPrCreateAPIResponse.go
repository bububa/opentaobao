package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurPrCreateAPIResponse 下pr单 API返回值
// alibaba.pur.pr.create
//
// 下pr单
type AlibabaPurPrCreateAPIResponse struct {
	model.CommonResponse
	AlibabaPurPrCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurPrCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurPrCreateAPIResponseModel).Reset()
}

// AlibabaPurPrCreateAPIResponseModel is 下pr单 成功返回结果
type AlibabaPurPrCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_pr_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回对象
	Result *MallReceivePrResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurPrCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurPrCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurPrCreateAPIResponse)
	},
}

// GetAlibabaPurPrCreateAPIResponse 从 sync.Pool 获取 AlibabaPurPrCreateAPIResponse
func GetAlibabaPurPrCreateAPIResponse() *AlibabaPurPrCreateAPIResponse {
	return poolAlibabaPurPrCreateAPIResponse.Get().(*AlibabaPurPrCreateAPIResponse)
}

// ReleaseAlibabaPurPrCreateAPIResponse 将 AlibabaPurPrCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurPrCreateAPIResponse(v *AlibabaPurPrCreateAPIResponse) {
	v.Reset()
	poolAlibabaPurPrCreateAPIResponse.Put(v)
}
