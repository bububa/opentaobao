package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建初始模板 APIResponse
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板
*/
type AlibabaWholesaleShippinglineTemplateInitAPIResponse struct {
    model.CommonResponse
    AlibabaWholesaleShippinglineTemplateInitResponse
}

type AlibabaWholesaleShippinglineTemplateInitResponse struct {
    XMLName xml.Name `xml:"alibaba_wholesale_shippingline_template_init_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 模板是否创建成功
    
    InitTemplateResponse   bool `json:"init_template_response,omitempty" xml:"init_template_response,omitempty"`

    
}
