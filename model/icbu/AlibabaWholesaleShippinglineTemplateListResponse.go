package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取运费模板 APIResponse
alibaba.wholesale.shippingline.template.list

查询运费模板信息
*/
type AlibabaWholesaleShippinglineTemplateListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWholesaleShippinglineTemplateListResponse `json:"alibaba_wholesale_shippingline_template_list_response,omitempty"`
}

type AlibabaWholesaleShippinglineTemplateListResponse struct {

    // 运费模板列表
    ListTemplateResponse   *ListTemplateAPIResult `json:"list_template_response,omitempty"`

}
