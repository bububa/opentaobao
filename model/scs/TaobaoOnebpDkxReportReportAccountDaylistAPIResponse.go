package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportAccountDaylistAPIResponse 获取账户分日报表 API返回值
// taobao.onebp.dkx.report.report.account.daylist
//
// 获取账户分日报表
type TaobaoOnebpDkxReportReportAccountDaylistAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportAccountDaylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel is 获取账户分日报表 成功返回结果
type TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_account_daylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportAccountDaylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportAccountDaylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportAccountDaylistAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportAccountDaylistAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportAccountDaylistAPIResponse
func GetTaobaoOnebpDkxReportReportAccountDaylistAPIResponse() *TaobaoOnebpDkxReportReportAccountDaylistAPIResponse {
	return poolTaobaoOnebpDkxReportReportAccountDaylistAPIResponse.Get().(*TaobaoOnebpDkxReportReportAccountDaylistAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportAccountDaylistAPIResponse 将 TaobaoOnebpDkxReportReportAccountDaylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportAccountDaylistAPIResponse(v *TaobaoOnebpDkxReportReportAccountDaylistAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportAccountDaylistAPIResponse.Put(v)
}
