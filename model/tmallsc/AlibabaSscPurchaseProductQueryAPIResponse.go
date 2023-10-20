package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscPurchaseProductQueryAPIResponse 查询已采购的服务产品 API返回值
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
type AlibabaSscPurchaseProductQueryAPIResponse struct {
	model.CommonResponse
	AlibabaSscPurchaseProductQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscPurchaseProductQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscPurchaseProductQueryAPIResponseModel).Reset()
}

// AlibabaSscPurchaseProductQueryAPIResponseModel is 查询已采购的服务产品 成功返回结果
type AlibabaSscPurchaseProductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_purchase_product_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务产品list
	ServiceProducts []ServiceProduct `json:"service_products,omitempty" xml:"service_products>service_product,omitempty"`
	// 错误展示信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscPurchaseProductQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceProducts = m.ServiceProducts[:0]
	m.DisplayMsg = ""
	m.IsSuccess = false
}

var poolAlibabaSscPurchaseProductQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscPurchaseProductQueryAPIResponse)
	},
}

// GetAlibabaSscPurchaseProductQueryAPIResponse 从 sync.Pool 获取 AlibabaSscPurchaseProductQueryAPIResponse
func GetAlibabaSscPurchaseProductQueryAPIResponse() *AlibabaSscPurchaseProductQueryAPIResponse {
	return poolAlibabaSscPurchaseProductQueryAPIResponse.Get().(*AlibabaSscPurchaseProductQueryAPIResponse)
}

// ReleaseAlibabaSscPurchaseProductQueryAPIResponse 将 AlibabaSscPurchaseProductQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscPurchaseProductQueryAPIResponse(v *AlibabaSscPurchaseProductQueryAPIResponse) {
	v.Reset()
	poolAlibabaSscPurchaseProductQueryAPIResponse.Put(v)
}
