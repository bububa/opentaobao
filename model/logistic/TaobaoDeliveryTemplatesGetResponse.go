package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户下所有模板 APIResponse
taobao.delivery.templates.get

根据用户ID获取用户下所有模板
*/
type TaobaoDeliveryTemplatesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeliveryTemplatesGetResponse `json:"taobao_delivery_templates_get_response,omitempty"`
}

type TaobaoDeliveryTemplatesGetResponse struct {

    // 运费模板列表
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty"`

    // 获得到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
