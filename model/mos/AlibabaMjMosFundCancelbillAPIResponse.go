package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundCancelbillAPIResponse 取消付款单 API返回值
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
type AlibabaMjMosFundCancelbillAPIResponse struct {
	model.CommonResponse
	AlibabaMjMosFundCancelbillAPIResponseModel
}

// AlibabaMjMosFundCancelbillAPIResponseModel is 取消付款单 成功返回结果
type AlibabaMjMosFundCancelbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_cancelbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
