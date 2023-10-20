package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderStatusQueryAPIResponse 查询电商券履约单状态 API返回值
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
type TmallAlihouseTradeCouponOrderStatusQueryAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponOrderStatusQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderStatusQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAlihouseTradeCouponOrderStatusQueryAPIResponseModel).Reset()
}

// TmallAlihouseTradeCouponOrderStatusQueryAPIResponseModel is 查询电商券履约单状态 成功返回结果
type TmallAlihouseTradeCouponOrderStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAlihouseTradeCouponOrderStatusQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderStatusQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAlihouseTradeCouponOrderStatusQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderStatusQueryAPIResponse)
	},
}

// GetTmallAlihouseTradeCouponOrderStatusQueryAPIResponse 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderStatusQueryAPIResponse
func GetTmallAlihouseTradeCouponOrderStatusQueryAPIResponse() *TmallAlihouseTradeCouponOrderStatusQueryAPIResponse {
	return poolTmallAlihouseTradeCouponOrderStatusQueryAPIResponse.Get().(*TmallAlihouseTradeCouponOrderStatusQueryAPIResponse)
}

// ReleaseTmallAlihouseTradeCouponOrderStatusQueryAPIResponse 将 TmallAlihouseTradeCouponOrderStatusQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderStatusQueryAPIResponse(v *TmallAlihouseTradeCouponOrderStatusQueryAPIResponse) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderStatusQueryAPIResponse.Put(v)
}
