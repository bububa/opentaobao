package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListAPIResponse 获取人群离线报表 API返回值
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;effect&#34;:15,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;crowd_id&#34;:12297883}
type TaobaoOnebpDkxReportReportCrowdListAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCrowdListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCrowdListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportCrowdListAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportCrowdListAPIResponseModel is 获取人群离线报表 成功返回结果
type TaobaoOnebpDkxReportReportCrowdListAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_crowd_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCrowdListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCrowdListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportCrowdListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportCrowdListAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportCrowdListAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportCrowdListAPIResponse
func GetTaobaoOnebpDkxReportReportCrowdListAPIResponse() *TaobaoOnebpDkxReportReportCrowdListAPIResponse {
	return poolTaobaoOnebpDkxReportReportCrowdListAPIResponse.Get().(*TaobaoOnebpDkxReportReportCrowdListAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportCrowdListAPIResponse 将 TaobaoOnebpDkxReportReportCrowdListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportCrowdListAPIResponse(v *TaobaoOnebpDkxReportReportCrowdListAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportCrowdListAPIResponse.Put(v)
}
