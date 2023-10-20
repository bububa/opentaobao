package icbuseller

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorOrderListAPIResponse 国际站服务市场订单列表接口 API返回值
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
type AlibabaSellerVendorOrderListAPIResponse struct {
	model.CommonResponse
	AlibabaSellerVendorOrderListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSellerVendorOrderListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSellerVendorOrderListAPIResponseModel).Reset()
}

// AlibabaSellerVendorOrderListAPIResponseModel is 国际站服务市场订单列表接口 成功返回结果
type AlibabaSellerVendorOrderListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaSellerVendorOrderListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSellerVendorOrderListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSellerVendorOrderListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorOrderListAPIResponse)
	},
}

// GetAlibabaSellerVendorOrderListAPIResponse 从 sync.Pool 获取 AlibabaSellerVendorOrderListAPIResponse
func GetAlibabaSellerVendorOrderListAPIResponse() *AlibabaSellerVendorOrderListAPIResponse {
	return poolAlibabaSellerVendorOrderListAPIResponse.Get().(*AlibabaSellerVendorOrderListAPIResponse)
}

// ReleaseAlibabaSellerVendorOrderListAPIResponse 将 AlibabaSellerVendorOrderListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSellerVendorOrderListAPIResponse(v *AlibabaSellerVendorOrderListAPIResponse) {
	v.Reset()
	poolAlibabaSellerVendorOrderListAPIResponse.Put(v)
}
