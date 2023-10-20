package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurprcreateAPIResponse 下pr单 API返回值
// alibaba.pur.pr.create
//
// 下pr单
type AlibabapurprcreateAPIResponse struct {
	model.CommonResponse
	AlibabapurprcreateAPIResponseModel
}

// AlibabapurprcreateAPIResponseModel is 下pr单 成功返回结果
type AlibabapurprcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_pr_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回对象
	Result *MallReceivePrResponse `json:"result,omitempty" xml:"result,omitempty"`
}
