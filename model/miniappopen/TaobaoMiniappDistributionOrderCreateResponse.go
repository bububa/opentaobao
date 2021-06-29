package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建小程序投放计划 APIResponse
taobao.miniapp.distribution.order.create

帮助商家，创建小程序的投放计划。
*/
type TaobaoMiniappDistributionOrderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappDistributionOrderCreateResponse
}

type TaobaoMiniappDistributionOrderCreateResponse struct {
    XMLName xml.Name `xml:"miniapp_distribution_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功
    
    OrderSuccess   bool `json:"order_success,omitempty" xml:"order_success,omitempty"`

    
    // 错误码
    
    OrderErrorCode   int64 `json:"order_error_code,omitempty" xml:"order_error_code,omitempty"`

    
    // 错误信息
    
    OrderErrorInfo   string `json:"order_error_info,omitempty" xml:"order_error_info,omitempty"`

    
    // 创建成功的投放计划id
    
    Model   int64 `json:"model,omitempty" xml:"model,omitempty"`

    
}
