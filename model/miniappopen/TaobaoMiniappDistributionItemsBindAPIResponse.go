package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionItemsBindAPIResponse
小程序投放-商品绑定/解绑 API返回值
taobao.miniapp.distribution.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。 */
type TaobaoMiniappDistributionItemsBindAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionItemsBindAPIResponseModel
}

// TaobaoMiniappDistributionItemsBindAPIResponseModel is 小程序投放-商品绑定/解绑 成功返回结果
type TaobaoMiniappDistributionItemsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定的结果列表
	ModelList []DistributionOrderBindTargetEntityOpenResultDto `json:"model_list,omitempty" xml:"model_list>distribution_order_bind_target_entity_open_result_dto,omitempty"`
}
