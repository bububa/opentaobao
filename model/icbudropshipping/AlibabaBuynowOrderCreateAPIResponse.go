package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBuynowOrderCreateAPIResponse 阿里巴巴买家buynow下单接口 API返回值
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
type AlibabaBuynowOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaBuynowOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBuynowOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBuynowOrderCreateAPIResponseModel).Reset()
}

// AlibabaBuynowOrderCreateAPIResponseModel is 阿里巴巴买家buynow下单接口 成功返回结果
type AlibabaBuynowOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_buynow_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Response
	Value *OrderCreateResponse `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBuynowOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = nil
}

var poolAlibabaBuynowOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBuynowOrderCreateAPIResponse)
	},
}

// GetAlibabaBuynowOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaBuynowOrderCreateAPIResponse
func GetAlibabaBuynowOrderCreateAPIResponse() *AlibabaBuynowOrderCreateAPIResponse {
	return poolAlibabaBuynowOrderCreateAPIResponse.Get().(*AlibabaBuynowOrderCreateAPIResponse)
}

// ReleaseAlibabaBuynowOrderCreateAPIResponse 将 AlibabaBuynowOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBuynowOrderCreateAPIResponse(v *AlibabaBuynowOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaBuynowOrderCreateAPIResponse.Put(v)
}
