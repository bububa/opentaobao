package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointcollectionstandpointqueryAPIResponse 查询收藏口径 API返回值
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
type AlibabalegalstandpointcollectionstandpointqueryAPIResponse struct {
	model.CommonResponse
	AlibabalegalstandpointcollectionstandpointqueryAPIResponseModel
}

// AlibabalegalstandpointcollectionstandpointqueryAPIResponseModel is 查询收藏口径 成功返回结果
type AlibabalegalstandpointcollectionstandpointqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_collectionstandpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 500
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 收藏口径
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// true
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
