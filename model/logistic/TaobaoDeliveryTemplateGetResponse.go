package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户指定运费模板信息 APIResponse
taobao.delivery.template.get

获取用户指定运费模板信息
*/
type TaobaoDeliveryTemplateGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoDeliveryTemplateGetResponse `json:"delivery_template_get_response,omitempty"` 
    TaobaoDeliveryTemplateGetResponse
}

/* model for simplify = false
type TaobaoDeliveryTemplateGetResponse struct {

    // 运费模板列表
    
    DeliveryTemplates  struct {
        DeliveryTemplate  []DeliveryTemplate `json:"delivery_template,omitempty"`
    } `json:"delivery_templates,omitempty"`
    

    // 获得到符合条件的结果总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoDeliveryTemplateGetResponse struct {

    // 运费模板列表
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty"`

    // 获得到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
