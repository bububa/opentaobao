package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointCollectionAPIResponse 收藏|取消收藏 API返回值
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
type AlibabaLegalStandpointStandpointCollectionAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointStandpointCollectionAPIResponseModel
}

// AlibabaLegalStandpointStandpointCollectionAPIResponseModel is 收藏|取消收藏 成功返回结果
type AlibabaLegalStandpointStandpointCollectionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpoint_collection_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 500
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// true
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
	// true
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
