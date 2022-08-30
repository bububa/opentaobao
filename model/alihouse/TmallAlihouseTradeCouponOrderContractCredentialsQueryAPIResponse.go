package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponse 查询用于电商券履约单合同下载的临时访问凭证 API返回值
// tmall.alihouse.trade.coupon.order.contract.credentials.query
//
// 获取用于下载合同的临时aksk和安全token
type TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponseModel
}

// TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponseModel is 查询用于电商券履约单合同下载的临时访问凭证 成功返回结果
type TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_contract_credentials_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TmallAlihouseTradeCouponOrderContractCredentialsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
