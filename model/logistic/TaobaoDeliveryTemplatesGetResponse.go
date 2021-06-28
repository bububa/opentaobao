package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户下所有模板 APIResponse
taobao.delivery.templates.get

根据用户ID获取用户下所有模板
*/
type TaobaoDeliveryTemplatesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"delivery_templates_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 运费模板列表
    
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty" xml:"