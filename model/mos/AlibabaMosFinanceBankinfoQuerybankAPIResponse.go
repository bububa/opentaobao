package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosFinanceBankinfoQuerybankAPIResponse 供应商银行账号查询 API返回值
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
type AlibabaMosFinanceBankinfoQuerybankAPIResponse struct {
	model.CommonResponse
	AlibabaMosFinanceBankinfoQuerybankAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosFinanceBankinfoQuerybankAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosFinanceBankinfoQuerybankAPIResponseModel).Reset()
}

// AlibabaMosFinanceBankinfoQuerybankAPIResponseModel is 供应商银行账号查询 成功返回结果
type AlibabaMosFinanceBankinfoQuerybankAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_finance_bankinfo_querybank_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosFinanceBankinfoQuerybankResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosFinanceBankinfoQuerybankAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosFinanceBankinfoQuerybankAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosFinanceBankinfoQuerybankAPIResponse)
	},
}

// GetAlibabaMosFinanceBankinfoQuerybankAPIResponse 从 sync.Pool 获取 AlibabaMosFinanceBankinfoQuerybankAPIResponse
func GetAlibabaMosFinanceBankinfoQuerybankAPIResponse() *AlibabaMosFinanceBankinfoQuerybankAPIResponse {
	return poolAlibabaMosFinanceBankinfoQuerybankAPIResponse.Get().(*AlibabaMosFinanceBankinfoQuerybankAPIResponse)
}

// ReleaseAlibabaMosFinanceBankinfoQuerybankAPIResponse 将 AlibabaMosFinanceBankinfoQuerybankAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosFinanceBankinfoQuerybankAPIResponse(v *AlibabaMosFinanceBankinfoQuerybankAPIResponse) {
	v.Reset()
	poolAlibabaMosFinanceBankinfoQuerybankAPIResponse.Put(v)
}
