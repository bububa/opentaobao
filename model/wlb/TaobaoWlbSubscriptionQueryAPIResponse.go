package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbSubscriptionQueryAPIResponse
查询商家定购的所有服务 API返回值
taobao.wlb.subscription.query

查询商家定购的所有服务,可通过入参状态来筛选 */
type TaobaoWlbSubscriptionQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbSubscriptionQueryAPIResponseModel
}

// TaobaoWlbSubscriptionQueryAPIResponseModel is 查询商家定购的所有服务 成功返回结果
type TaobaoWlbSubscriptionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_subscription_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 卖家定购的服务列表
	SellerSubscriptionList []WlbSellerSubscription `json:"seller_subscription_list,omitempty" xml:"seller_subscription_list>wlb_seller_subscription,omitempty"`
}
