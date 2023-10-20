package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaHourReportAccountGetAPIResponse 账户级别小时报表获取 API返回值
// taobao.simba.hour.report.account.get
//
// 获取账户小时实时报表数据
type TaobaoSimbaHourReportAccountGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaHourReportAccountGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaHourReportAccountGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaHourReportAccountGetAPIResponseModel).Reset()
}

// TaobaoSimbaHourReportAccountGetAPIResponseModel is 账户级别小时报表获取 成功返回结果
type TaobaoSimbaHourReportAccountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_hour_report_account_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaHourReportAccountGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaHourReportAccountGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaHourReportAccountGetAPIResponse)
	},
}

// GetTaobaoSimbaHourReportAccountGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaHourReportAccountGetAPIResponse
func GetTaobaoSimbaHourReportAccountGetAPIResponse() *TaobaoSimbaHourReportAccountGetAPIResponse {
	return poolTaobaoSimbaHourReportAccountGetAPIResponse.Get().(*TaobaoSimbaHourReportAccountGetAPIResponse)
}

// ReleaseTaobaoSimbaHourReportAccountGetAPIResponse 将 TaobaoSimbaHourReportAccountGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaHourReportAccountGetAPIResponse(v *TaobaoSimbaHourReportAccountGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaHourReportAccountGetAPIResponse.Put(v)
}
