package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmosfundcreatebillAPIResponse 创建一个付款单 API返回值
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
type AlibabamjmosfundcreatebillAPIResponse struct {
	model.CommonResponse
	AlibabamjmosfundcreatebillAPIResponseModel
}

// AlibabamjmosfundcreatebillAPIResponseModel is 创建一个付款单 成功返回结果
type AlibabamjmosfundcreatebillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_createbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
