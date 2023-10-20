package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportAccountOfflineAPIResponse 获取账户历史报表 API返回值
// taobao.onebp.dkx.report.report.account.offline
//
// 获取账户历史报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;start_time&#34;: &#34;2021-07-24&#34;,     &#34;effect&#34;: 15,     &#34;end_time&#34;: &#34;2021-08-21&#34;,     &#34;strategy_scene&#34;:true,     &#34;unify_type&#34;:&#34;kuan&#34;,     &#34;bizCode&#34;:&#34;adStrategyDkx&#34;   }
type TaobaoOnebpDkxReportReportAccountOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportAccountOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportAccountOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportAccountOfflineAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportAccountOfflineAPIResponseModel is 获取账户历史报表 成功返回结果
type TaobaoOnebpDkxReportReportAccountOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_account_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportAccountOfflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportAccountOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportAccountOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportAccountOfflineAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportAccountOfflineAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportAccountOfflineAPIResponse
func GetTaobaoOnebpDkxReportReportAccountOfflineAPIResponse() *TaobaoOnebpDkxReportReportAccountOfflineAPIResponse {
	return poolTaobaoOnebpDkxReportReportAccountOfflineAPIResponse.Get().(*TaobaoOnebpDkxReportReportAccountOfflineAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportAccountOfflineAPIResponse 将 TaobaoOnebpDkxReportReportAccountOfflineAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportAccountOfflineAPIResponse(v *TaobaoOnebpDkxReportReportAccountOfflineAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportAccountOfflineAPIResponse.Put(v)
}
