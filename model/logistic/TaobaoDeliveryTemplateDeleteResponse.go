package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除运费模板 APIResponse
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
type TaobaoDeliveryTemplateDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeliveryTemplateDeleteResponse `json:"taobao_delivery_template_delete_response,omitempty"`
}

type TaobaoDeliveryTemplateDeleteResponse struct {

    // 表示删除成功还是失败
    Complete   bool `json:"complete,omitempty"`

}
