package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialRuleUpdateAPIResponse 专属下单更新限购规则 API返回值
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
type TaobaoOpentradeSpecialRuleUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeSpecialRuleUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialRuleUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeSpecialRuleUpdateAPIResponseModel).Reset()
}

// TaobaoOpentradeSpecialRuleUpdateAPIResponseModel is 专属下单更新限购规则 成功返回结果
type TaobaoOpentradeSpecialRuleUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_special_rule_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新失败的商品列表
	Result []ItemResultDto `json:"result,omitempty" xml:"result>item_result_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialRuleUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
}

var poolTaobaoOpentradeSpecialRuleUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeSpecialRuleUpdateAPIResponse)
	},
}

// GetTaobaoOpentradeSpecialRuleUpdateAPIResponse 从 sync.Pool 获取 TaobaoOpentradeSpecialRuleUpdateAPIResponse
func GetTaobaoOpentradeSpecialRuleUpdateAPIResponse() *TaobaoOpentradeSpecialRuleUpdateAPIResponse {
	return poolTaobaoOpentradeSpecialRuleUpdateAPIResponse.Get().(*TaobaoOpentradeSpecialRuleUpdateAPIResponse)
}

// ReleaseTaobaoOpentradeSpecialRuleUpdateAPIResponse 将 TaobaoOpentradeSpecialRuleUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeSpecialRuleUpdateAPIResponse(v *TaobaoOpentradeSpecialRuleUpdateAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeSpecialRuleUpdateAPIResponse.Put(v)
}
