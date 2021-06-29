package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条新增 APIResponse
taobao.alitrip.it.policy.add

销售规则新增，成功返回taobaoId
*/
type TaobaoAlitripItPolicyAddAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItPolicyAddResponse
}

type TaobaoAlitripItPolicyAddResponse struct {
    XMLName xml.Name `xml:"alitrip_it_policy_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 淘宝政策id
    
    TaobaoId   int64 `json:"taobao_id,omitempty" xml:"taobao_id,omitempty"`

    
    // 扩展字段
    
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`

    
}
