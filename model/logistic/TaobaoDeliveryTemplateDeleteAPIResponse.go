package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除运费模板 API返回值 
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
type TaobaoDeliveryTemplateDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoDeliveryTemplateDeleteAPIResponseModel
}

// 删除运费模板 成功返回结果
type TaobaoDeliveryTemplateDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"delivery_template_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 表示删除成功还是失败
    Complete   bool `json:"complete,omitempty" xml:"complete,omitempty"`
}
