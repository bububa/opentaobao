package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderListQueryAPIResponse 星河-订单列表查询 API返回值
// alitrip.merchant.galaxy.order.list.query
//
// 为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
type AlitripMerchantGalaxyOrderListQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderListQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderListQueryAPIResponseModel is 星河-订单列表查询 成功返回结果
type AlitripMerchantGalaxyOrderListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderListQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderListQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderListQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderListQueryAPIResponse
func GetAlitripMerchantGalaxyOrderListQueryAPIResponse() *AlitripMerchantGalaxyOrderListQueryAPIResponse {
	return poolAlitripMerchantGalaxyOrderListQueryAPIResponse.Get().(*AlitripMerchantGalaxyOrderListQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderListQueryAPIResponse 将 AlitripMerchantGalaxyOrderListQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderListQueryAPIResponse(v *AlitripMerchantGalaxyOrderListQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderListQueryAPIResponse.Put(v)
}
