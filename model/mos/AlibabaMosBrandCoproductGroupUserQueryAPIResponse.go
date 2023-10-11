package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosBrandCoproductGroupUserQueryAPIResponse 按照条件查询分页数据 API返回值
// alibaba.mos.brand.coproduct.group.user.query
//
// 按照条件查询分页数据
type AlibabaMosBrandCoproductGroupUserQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMosBrandCoproductGroupUserQueryAPIResponseModel
}

// AlibabaMosBrandCoproductGroupUserQueryAPIResponseModel is 按照条件查询分页数据 成功返回结果
type AlibabaMosBrandCoproductGroupUserQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_brand_coproduct_group_user_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 组员列表信息
	Data []BrandCoProductGroupUserDto `json:"data,omitempty" xml:"data>brand_co_product_group_user_dto,omitempty"`
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 异步结果
	AsyncResult string `json:"async_result,omitempty" xml:"async_result,omitempty"`
	// 错误code
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 调用是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// 是否同步
	IsAsync bool `json:"is_async,omitempty" xml:"is_async,omitempty"`
}
