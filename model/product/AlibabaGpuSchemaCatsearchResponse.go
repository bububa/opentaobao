package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按类目查询spu接口 APIResponse
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口
*/
type AlibabaGpuSchemaCatsearchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaGpuSchemaCatsearchResponse `json:"alibaba_gpu_schema_catsearch_response,omitempty"` 
    AlibabaGpuSchemaCatsearchResponse
}

/* model for simplify = false
type AlibabaGpuSchemaCatsearchResponse struct {

    // 返回按类目查询spu的schema
    
    CatSearchResult   string `json:"cat_search_result,omitempty"`
    

    // 总记录数
    
    Total   int64 `json:"total,omitempty"`
    

}
*/

type AlibabaGpuSchemaCatsearchResponse struct {

    // 返回按类目查询spu的schema
    CatSearchResult   string `json:"cat_search_result,omitempty"`

    // 总记录数
    Total   int64 `json:"total,omitempty"`

}
