package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)层级属性获取 APIResponse
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回
*/
type AlibabaIcbuCategorySchemaLevelGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategorySchemaLevelGetResponse `json:"alibaba_icbu_category_schema_level_get_response,omitempty"` 
    AlibabaIcbuCategorySchemaLevelGetResponse
}

/* model for simplify = false
type AlibabaIcbuCategorySchemaLevelGetResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaIcbuCategorySchemaLevelGetResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
