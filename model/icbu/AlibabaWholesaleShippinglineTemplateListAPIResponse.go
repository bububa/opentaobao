package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运费模板 API返回值 
alibaba.wholesale.shippingline.template.list

查询运费模板信息
*/
type AlibabaWholesaleShippinglineTemplateListAPIResponse struct {
    model.CommonResponse
    AlibabaWholesaleShippinglineTemplateListAPIResponseModel
}

// 获取运费模板 成功返回结果
type AlibabaWholesaleShippinglineTemplateListAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wholesale_shippingline_template_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 运费模板列表
    ListTemplateResponse   *ListTemplateAPIResult `json:"list_template_response,omitempty" xml:"list_template_response,omitempty"`
}
