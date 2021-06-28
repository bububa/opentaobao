package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改运费模板 APIResponse
taobao.delivery.template.update

修改运费模板
*/
type TaobaoDeliveryTemplateUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoDeliveryTemplateUpdateResponse
}

type TaobaoDeliveryTemplateUpdateResponse struct {
    XMLName xml.Name `xml:"delivery_template_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 表示修改是否成功
    
    Complete   bool `json:"complete,omitempty" xml:"complete,omitempty"`

    
}
