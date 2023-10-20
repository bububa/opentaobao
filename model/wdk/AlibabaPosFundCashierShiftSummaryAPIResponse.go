package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPosFundCashierShiftSummaryAPIResponse 收银换班数据同步 API返回值
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
type AlibabaPosFundCashierShiftSummaryAPIResponse struct {
	model.CommonResponse
	AlibabaPosFundCashierShiftSummaryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPosFundCashierShiftSummaryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPosFundCashierShiftSummaryAPIResponseModel).Reset()
}

// AlibabaPosFundCashierShiftSummaryAPIResponseModel is 收银换班数据同步 成功返回结果
type AlibabaPosFundCashierShiftSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pos_fund_cashier_shift_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaPosFundCashierShiftSummaryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPosFundCashierShiftSummaryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPosFundCashierShiftSummaryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPosFundCashierShiftSummaryAPIResponse)
	},
}

// GetAlibabaPosFundCashierShiftSummaryAPIResponse 从 sync.Pool 获取 AlibabaPosFundCashierShiftSummaryAPIResponse
func GetAlibabaPosFundCashierShiftSummaryAPIResponse() *AlibabaPosFundCashierShiftSummaryAPIResponse {
	return poolAlibabaPosFundCashierShiftSummaryAPIResponse.Get().(*AlibabaPosFundCashierShiftSummaryAPIResponse)
}

// ReleaseAlibabaPosFundCashierShiftSummaryAPIResponse 将 AlibabaPosFundCashierShiftSummaryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPosFundCashierShiftSummaryAPIResponse(v *AlibabaPosFundCashierShiftSummaryAPIResponse) {
	v.Reset()
	poolAlibabaPosFundCashierShiftSummaryAPIResponse.Put(v)
}
