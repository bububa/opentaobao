package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增运费模板 APIResponse
taobao.delivery.template.add

增加运费模板的外部接口
*/
type TaobaoDeliveryTemplateAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoDeliveryTemplateAddResponse `json:"delivery_template_add_response,omitempty"` 
    TaobaoDeliveryTemplateAddResponse
}

/* model for simplify = false
type TaobaoDeliveryTemplateAddResponse struct {

    // 模板对象
    
    DeliveryTemplate  *struct {
        DeliveryTemplate  *DeliveryTemplate `json:"delivery_template,omitempty"`
    } `json:"delivery_template,omitempty"`
    

}
*/

type TaobaoDeliveryTemplateAddResponse struct {

    // 模板对象
    DeliveryTemplate   *DeliveryTemplate `json:"delivery_template,omitempty"`

}
