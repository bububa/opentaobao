package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponordercontractkeyqueryAPIResponse 查询电商券履约单合同key API返回值
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
type TmallalihousetradecouponordercontractkeyqueryAPIResponse struct {
	model.CommonResponse
	TmallalihousetradecouponordercontractkeyqueryAPIResponseModel
}

// TmallalihousetradecouponordercontractkeyqueryAPIResponseModel is 查询电商券履约单合同key 成功返回结果
type TmallalihousetradecouponordercontractkeyqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_contract_key_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallalihousetradecouponordercontractkeyqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
