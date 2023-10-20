package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmosfundmodifybillbankaccountAPIResponse 修改付款单的银行账户信息 API返回值
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
type AlibabamjmosfundmodifybillbankaccountAPIResponse struct {
	model.CommonResponse
	AlibabamjmosfundmodifybillbankaccountAPIResponseModel
}

// AlibabamjmosfundmodifybillbankaccountAPIResponseModel is 修改付款单的银行账户信息 成功返回结果
type AlibabamjmosfundmodifybillbankaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_modifybillbankaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
