package scs

import (
	"encoding/xml"

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

// TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel is 获取账户分日报表 成功返回结果
type TaobaoOnebpDkxReportReportAccountDaylistAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_account_daylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportAccountDaylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
