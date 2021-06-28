package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新旧属性的映射 APIResponse
alibaba.icbu.category.id.mapping

商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
*/
type AlibabaIcbuCategoryIdMappingAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuCategoryIdMappingResponse `json:"alibaba_icbu_category_id_mapping_response,omitempty"` 
    AlibabaIcbuCategoryIdMappingResponse
}

/* model for simplify = false
type AlibabaIcbuCategoryIdMappingResponse struct {

    // 转化的类目id
    
    MappingResult   int64 `json:"mapping_result,omitempty"`
    

}
*/

type AlibabaIcbuCategoryIdMappingResponse struct {

    // 转化的类目id
    MappingResult   int64 `json:"mapping_result,omitempty"`

}
