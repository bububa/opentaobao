package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户下所有模板 API返回值 
taobao.delivery.templates.get

根据用户ID获取用户下所有模板
*/
type TaobaoDeliveryTemplatesGetAPIResponse struct {
    model.CommonResponse
    TaobaoDeliveryTemplatesGetAPIResponseModel
}

// 获取用户下所有模板 成功返回结果
type TaobaoDeliveryTemplatesGetAPIResponseModel struct {
    XMLName xml.Name `xml:"delivery_templates_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 运费模板列表
    DeliveryTemplates   []DeliveryTemplate `json:"delivery_templates,omitempty" xml:"delivery_templates>delivery_template,omitempty"`
    // 获得到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
