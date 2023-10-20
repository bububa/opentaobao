package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMosFundModifybillbankaccountAPIResponse 修改付款单的银行账户信息 API返回值
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
type AlibabaMjMosFundModifybillbankaccountAPIResponse struct {
	model.CommonResponse
	AlibabaMjMosFundModifybillbankaccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMosFundModifybillbankaccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMosFundModifybillbankaccountAPIResponseModel).Reset()
}

// AlibabaMjMosFundModifybillbankaccountAPIResponseModel is 修改付款单的银行账户信息 成功返回结果
type AlibabaMjMosFundModifybillbankaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_mos_fund_modifybillbankaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMosFundModifybillbankaccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolAlibabaMjMosFundModifybillbankaccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMosFundModifybillbankaccountAPIResponse)
	},
}

// GetAlibabaMjMosFundModifybillbankaccountAPIResponse 从 sync.Pool 获取 AlibabaMjMosFundModifybillbankaccountAPIResponse
func GetAlibabaMjMosFundModifybillbankaccountAPIResponse() *AlibabaMjMosFundModifybillbankaccountAPIResponse {
	return poolAlibabaMjMosFundModifybillbankaccountAPIResponse.Get().(*AlibabaMjMosFundModifybillbankaccountAPIResponse)
}

// ReleaseAlibabaMjMosFundModifybillbankaccountAPIResponse 将 AlibabaMjMosFundModifybillbankaccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMosFundModifybillbankaccountAPIResponse(v *AlibabaMjMosFundModifybillbankaccountAPIResponse) {
	v.Reset()
	poolAlibabaMjMosFundModifybillbankaccountAPIResponse.Put(v)
}
