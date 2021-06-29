package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（已废弃）构建实例化应用 APIResponse
taobao.miniappp.template.instantiate

实例化saas化的小程序
*/
type TaobaoMiniapppTemplateInstantiateAPIResponse struct {
    model.CommonResponse
    TaobaoMiniapppTemplateInstantiateResponse
}

type TaobaoMiniapppTemplateInstantiateResponse struct {
    XMLName xml.Name `xml:"miniappp_template_instantiate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoMiniapppTemplateInstantiateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
