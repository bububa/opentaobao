package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderItemsBindAPIResponse 小程序投放-基于投放计划绑定/解绑商品 API返回值
// taobao.miniapp.distribution.order.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
type TaobaoMiniappDistributionOrderItemsBindAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionOrderItemsBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderItemsBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionOrderItemsBindAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionOrderItemsBindAPIResponseModel is 小程序投放-基于投放计划绑定/解绑商品 成功返回结果
type TaobaoMiniappDistributionOrderItemsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_order_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定的结果
	Model *DistributionOrderBindTargetEntityOpenResultDto `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderItemsBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolTaobaoMiniappDistributionOrderItemsBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionOrderItemsBindAPIResponse)
	},
}

// GetTaobaoMiniappDistributionOrderItemsBindAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionOrderItemsBindAPIResponse
func GetTaobaoMiniappDistributionOrderItemsBindAPIResponse() *TaobaoMiniappDistributionOrderItemsBindAPIResponse {
	return poolTaobaoMiniappDistributionOrderItemsBindAPIResponse.Get().(*TaobaoMiniappDistributionOrderItemsBindAPIResponse)
}

// ReleaseTaobaoMiniappDistributionOrderItemsBindAPIResponse 将 TaobaoMiniappDistributionOrderItemsBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderItemsBindAPIResponse(v *TaobaoMiniappDistributionOrderItemsBindAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderItemsBindAPIResponse.Put(v)
}
