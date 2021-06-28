package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建初始模板 APIResponse
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板
*/
type AlibabaWholesaleShippinglineTemplateInitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWholesaleShippinglineTemplateInitResponse `json:"alibaba_wholesale_shippingline_template_init_response,omitempty"` 
    AlibabaWholesaleShippinglineTemplateInitResponse
}

/* model for simplify = false
type AlibabaWholesaleShippinglineTemplateInitResponse struct {

    // 模板是否创建成功
    
    InitTemplateResponse   bool `json:"init_template_response,omitempty"`
    

}
*/

type AlibabaWholesaleShippinglineTemplateInitResponse struct {

    // 模板是否创建成功
    InitTemplateResponse   bool `json:"init_template_response,omitempty"`

}
