package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaposfundcashiershiftsummaryAPIResponse 收银换班数据同步 API返回值
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
type AlibabaposfundcashiershiftsummaryAPIResponse struct {
	model.CommonResponse
	AlibabaposfundcashiershiftsummaryAPIResponseModel
}

// AlibabaposfundcashiershiftsummaryAPIResponseModel is 收银换班数据同步 成功返回结果
type AlibabaposfundcashiershiftsummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pos_fund_cashier_shift_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaposfundcashiershiftsummaryResult `json:"result,omitempty" xml:"result,omitempty"`
}
