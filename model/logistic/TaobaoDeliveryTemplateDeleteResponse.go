package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除运费模板 APIResponse
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
type TaobaoDeliveryTemplateDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"delivery_template_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 表示删除成功还是失败
    
    Complete   bool `json:"complete,omitempty" xml:"