package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponorderstatusqueryAPIResponse 查询电商券履约单状态 API返回值
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
type TmallalihousetradecouponorderstatusqueryAPIResponse struct {
	model.CommonResponse
	TmallalihousetradecouponorderstatusqueryAPIResponseModel
}

// TmallalihousetradecouponorderstatusqueryAPIResponseModel is 查询电商券履约单状态 成功返回结果
type TmallalihousetradecouponorderstatusqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallalihousetradecouponorderstatusqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
