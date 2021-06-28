package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
层级属性的子属性获取 APIResponse
alibaba.icbu.category.level.attr.get

用于获取层级属性（车型库）的子属性和属性值
*/
type AlibabaIcbuCategoryLevelAttrGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_category_level_attr_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    ResultList   *AlibabaIcbuCategoryLevelAttrGetResult `json:"result_list,omitempty" xml:"