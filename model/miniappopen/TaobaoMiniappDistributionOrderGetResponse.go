package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-查询小程序投放计划信息 APIResponse
taobao.miniapp.distribution.order.get

服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
*/
type TaobaoMiniappDistributionOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappDistributionOrderGetResponse
}

type TaobaoMiniappDistributionOrderGetResponse struct {
    XMLName xml.Name `xml:"miniapp_distribution_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 投放计划列表的详细信息
    
    Model   []DistributionOrderOpenBizDto `json:"model,omitempty" xml:"model>distribution_order_open_biz_dto,omitempty"`
    
    
}
