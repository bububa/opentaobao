package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportAccountRealtimeAPIResponse 获取账户实时报表 API返回值
// taobao.onebp.dkx.report.report.account.realtime
//
// 获取账户实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;log_date_list&#34;: [  &#34;2021-09-23&#34;     ]   }
type TaobaoOnebpDkxReportReportAccountRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportAccountRealtimeAPIResponseModel
}

// TaobaoOnebpDkxReportReportAccountRealtimeAPIResponseModel is 获取账户实时报表 成功返回结果
type TaobaoOnebpDkxReportReportAccountRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_account_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportAccountRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
