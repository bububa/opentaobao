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
    Response *TaobaoDeliveryTemplateGetResponse `json:"taobao_delivery_template_get_response,omitempty"`
}

type TaobaoDeliveryTemplateGetResponse struct {

    // 运费模板列表
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty"`

    // 获得到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
