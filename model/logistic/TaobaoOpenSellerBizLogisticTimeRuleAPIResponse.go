package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenSellerBizLogisticTimeRuleAPIResponse 商家自定义发货时效 API返回值
// taobao.open.seller.biz.logistic.time.rule
//
// 服务商回传商家自定义发货时效
type TaobaoOpenSellerBizLogisticTimeRuleAPIResponse struct {
	model.CommonResponse
	TaobaoOpenSellerBizLogisticTimeRuleAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenSellerBizLogisticTimeRuleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenSellerBizLogisticTimeRuleAPIResponseModel).Reset()
}

// TaobaoOpenSellerBizLogisticTimeRuleAPIResponseModel is 商家自定义发货时效 成功返回结果
type TaobaoOpenSellerBizLogisticTimeRuleAPIResponseModel struct {
	XMLName xml.Name `xml:"open_seller_biz_logistic_time_rule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenSellerBizLogisticTimeRuleAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoOpenSellerBizLogisticTimeRuleAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenSellerBizLogisticTimeRuleAPIResponse)
	},
}

// GetTaobaoOpenSellerBizLogisticTimeRuleAPIResponse 从 sync.Pool 获取 TaobaoOpenSellerBizLogisticTimeRuleAPIResponse
func GetTaobaoOpenSellerBizLogisticTimeRuleAPIResponse() *TaobaoOpenSellerBizLogisticTimeRuleAPIResponse {
	return poolTaobaoOpenSellerBizLogisticTimeRuleAPIResponse.Get().(*TaobaoOpenSellerBizLogisticTimeRuleAPIResponse)
}

// ReleaseTaobaoOpenSellerBizLogisticTimeRuleAPIResponse 将 TaobaoOpenSellerBizLogisticTimeRuleAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenSellerBizLogisticTimeRuleAPIResponse(v *TaobaoOpenSellerBizLogisticTimeRuleAPIResponse) {
	v.Reset()
	poolTaobaoOpenSellerBizLogisticTimeRuleAPIResponse.Put(v)
}
