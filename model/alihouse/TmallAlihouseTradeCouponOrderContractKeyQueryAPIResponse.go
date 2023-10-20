package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse 查询电商券履约单合同key API返回值
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
type TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse struct {
	model.CommonResponse
	TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponseModel).Reset()
}

// TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponseModel is 查询电商券履约单合同key 成功返回结果
type TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_alihouse_trade_coupon_order_contract_key_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAlihouseTradeCouponOrderContractKeyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse)
	},
}

// GetTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse
func GetTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse() *TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse {
	return poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse.Get().(*TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse)
}

// ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse 将 TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse(v *TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse.Put(v)
}
