package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)层级属性获取 API返回值 
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回
*/
type AlibabaIcbuCategorySchemaLevelGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategorySchemaLevelGetResponse
}

// (新)层级属性获取 成功返回结果
type AlibabaIcbuCategorySchemaLevelGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_schema_level_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Top返回对象
    Result   *TopResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
