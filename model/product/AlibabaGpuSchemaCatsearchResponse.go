package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按类目查询spu接口 APIResponse
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口
*/
type AlibabaGpuSchemaCatsearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_gpu_schema_catsearch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回按类目查询spu的schema
    
    CatSearchResult   string `json:"cat_search_result,omitempty" xml:"