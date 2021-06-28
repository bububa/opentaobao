package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)层级属性获取 APIResponse
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回
*/
type AlibabaIcbuCategorySchemaLevelGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategorySchemaLevelGetResponse
}

type AlibabaIcbuCategorySchemaLevelGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_schema_level_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
