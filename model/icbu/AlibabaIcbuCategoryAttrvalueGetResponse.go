package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
属性值获取 APIResponse
alibaba.icbu.category.attrvalue.get

属性值获取
*/
type AlibabaIcbuCategoryAttrvalueGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_category_attrvalue_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    ResultList   []AttributeValue `json:"result_list,omitempty" xml:"