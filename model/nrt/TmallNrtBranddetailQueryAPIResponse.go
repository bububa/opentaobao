package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtbranddetailqueryAPIResponse 品牌详情查询 API返回值
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
type TmallnrtbranddetailqueryAPIResponse struct {
	model.CommonResponse
	TmallnrtbranddetailqueryAPIResponseModel
}

// TmallnrtbranddetailqueryAPIResponseModel is 品牌详情查询 成功返回结果
type TmallnrtbranddetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_branddetail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
