package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuSchemaCatsearchAPIResponse 按类目查询spu接口 API返回值
// alibaba.gpu.schema.catsearch
//
// 按类目查询spu的schema接口
type AlibabaGpuSchemaCatsearchAPIResponse struct {
	model.CommonResponse
	AlibabaGpuSchemaCatsearchAPIResponseModel
}

// AlibabaGpuSchemaCatsearchAPIResponseModel is 按类目查询spu接口 成功返回结果
type AlibabaGpuSchemaCatsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_schema_catsearch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回按类目查询spu的schema
	CatSearchResult string `json:"cat_search_result,omitempty" xml:"cat_search_result,omitempty"`
	// 总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
