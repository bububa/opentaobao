package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbSubscriptionQueryAPIResponse 查询商家定购的所有服务 API返回值
// taobao.wlb.subscription.query
//
// 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbSubscriptionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbSubscriptionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbSubscriptionQueryAPIResponseModel).Reset()
}

// TaobaoWlbSubscriptionQueryAPIResponseModel is 查询商家定购的所有服务 成功返回结果
type TaobaoWlbSubscriptionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_subscription_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖家定购的服务列表
	SellerSubscriptionList []SellerSubscriptionList `json:"seller_subscription_list,omitempty" xml:"seller_subscription_list>seller_subscription_list,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbSubscriptionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SellerSubscriptionList = m.SellerSubscriptionList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbSubscriptionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbSubscriptionQueryAPIResponse)
	},
}

// GetTaobaoWlbSubscriptionQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbSubscriptionQueryAPIResponse
func GetTaobaoWlbSubscriptionQueryAPIResponse() *TaobaoWlbSubscriptionQueryAPIResponse {
	return poolTaobaoWlbSubscriptionQueryAPIResponse.Get().(*TaobaoWlbSubscriptionQueryAPIResponse)
}

// ReleaseTaobaoWlbSubscriptionQueryAPIResponse 将 TaobaoWlbSubscriptionQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbSubscriptionQueryAPIResponse(v *TaobaoWlbSubscriptionQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbSubscriptionQueryAPIResponse.Put(v)
}
