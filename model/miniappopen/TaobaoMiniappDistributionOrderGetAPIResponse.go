package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderGetAPIResponse 小程序投放-查询小程序投放计划信息 API返回值
// taobao.miniapp.distribution.order.get
//
// 服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
type TaobaoMiniappDistributionOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionOrderGetAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionOrderGetAPIResponseModel is 小程序投放-查询小程序投放计划信息 成功返回结果
type TaobaoMiniappDistributionOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 投放计划列表的详细信息
	Model []DistributionOrderOpenBizDto `json:"model,omitempty" xml:"model>distribution_order_open_biz_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
}

var poolTaobaoMiniappDistributionOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionOrderGetAPIResponse)
	},
}

// GetTaobaoMiniappDistributionOrderGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionOrderGetAPIResponse
func GetTaobaoMiniappDistributionOrderGetAPIResponse() *TaobaoMiniappDistributionOrderGetAPIResponse {
	return poolTaobaoMiniappDistributionOrderGetAPIResponse.Get().(*TaobaoMiniappDistributionOrderGetAPIResponse)
}

// ReleaseTaobaoMiniappDistributionOrderGetAPIResponse 将 TaobaoMiniappDistributionOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderGetAPIResponse(v *TaobaoMiniappDistributionOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderGetAPIResponse.Put(v)
}
