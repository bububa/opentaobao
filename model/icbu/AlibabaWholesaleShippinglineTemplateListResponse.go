package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运费模板 APIResponse
alibaba.wholesale.shippingline.template.list

查询运费模板信息
*/
type AlibabaWholesaleShippinglineTemplateListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wholesale_shippingline_template_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 运费模板列表
    
    ListTemplateResponse   *ListTemplateAPIResult `json:"list_template_response,omitempty" xml:"