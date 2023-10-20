package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionItemsBindAPIResponse 【已废弃】小程序投放-商品绑定/解绑 API返回值
// taobao.miniapp.distribution.items.bind
//
// 【已废弃，请使用 taobao.miniapp.distribution.order.items.bind 接口】提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
type TaobaoMiniappDistributionItemsBindAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionItemsBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionItemsBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionItemsBindAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionItemsBindAPIResponseModel is 【已废弃】小程序投放-商品绑定/解绑 成功返回结果
type TaobaoMiniappDistributionItemsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定的结果列表
	ModelList []DistributionOrderBindTargetEntityOpenResultDto `json:"model_list,omitempty" xml:"model_list>distribution_order_bind_target_entity_open_result_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionItemsBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
}

var poolTaobaoMiniappDistributionItemsBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionItemsBindAPIResponse)
	},
}

// GetTaobaoMiniappDistributionItemsBindAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionItemsBindAPIResponse
func GetTaobaoMiniappDistributionItemsBindAPIResponse() *TaobaoMiniappDistributionItemsBindAPIResponse {
	return poolTaobaoMiniappDistributionItemsBindAPIResponse.Get().(*TaobaoMiniappDistributionItemsBindAPIResponse)
}

// ReleaseTaobaoMiniappDistributionItemsBindAPIResponse 将 TaobaoMiniappDistributionItemsBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionItemsBindAPIResponse(v *TaobaoMiniappDistributionItemsBindAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionItemsBindAPIResponse.Put(v)
}
