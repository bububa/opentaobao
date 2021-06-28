package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户指定运费模板信息 APIResponse
taobao.delivery.template.get

获取用户指定运费模板信息
*/
type TaobaoDeliveryTemplateGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"delivery_template_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 运费模板列表
    
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty" xml:"