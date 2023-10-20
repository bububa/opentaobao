package icbuseller

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorOrderDetailAPIResponse 国际站服务市场订单详情接口 API返回值
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
type AlibabaSellerVendorOrderDetailAPIResponse struct {
	model.CommonResponse
	AlibabaSellerVendorOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSellerVendorOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSellerVendorOrderDetailAPIResponseModel).Reset()
}

// AlibabaSellerVendorOrderDetailAPIResponseModel is 国际站服务市场订单详情接口 成功返回结果
type AlibabaSellerVendorOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回对象
	Result *AlibabaSellerVendorOrderDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSellerVendorOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSellerVendorOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorOrderDetailAPIResponse)
	},
}

// GetAlibabaSellerVendorOrderDetailAPIResponse 从 sync.Pool 获取 AlibabaSellerVendorOrderDetailAPIResponse
func GetAlibabaSellerVendorOrderDetailAPIResponse() *AlibabaSellerVendorOrderDetailAPIResponse {
	return poolAlibabaSellerVendorOrderDetailAPIResponse.Get().(*AlibabaSellerVendorOrderDetailAPIResponse)
}

// ReleaseAlibabaSellerVendorOrderDetailAPIResponse 将 AlibabaSellerVendorOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSellerVendorOrderDetailAPIResponse(v *AlibabaSellerVendorOrderDetailAPIResponse) {
	v.Reset()
	poolAlibabaSellerVendorOrderDetailAPIResponse.Put(v)
}
