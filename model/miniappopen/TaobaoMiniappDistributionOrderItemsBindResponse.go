package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-基于投放计划绑定/解绑商品 APIResponse
taobao.miniapp.distribution.order.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
*/
type TaobaoMiniappDistributionOrderItemsBindAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappDistributionOrderItemsBindResponse
}

type TaobaoMiniappDistributionOrderItemsBindResponse struct {
    XMLName xml.Name `xml:"miniapp_distribution_order_items_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 绑定的结果
    
    Model   *DistributionOrderBindTargetEntityOpenResultDTO `json:"model,omitempty" xml:"model,omitempty"`

    
}
