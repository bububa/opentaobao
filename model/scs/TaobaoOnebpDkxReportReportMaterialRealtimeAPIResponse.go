package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse 查询某计划分商品实时报表 API返回值
// taobao.onebp.dkx.report.report.material.realtime
//
// 查询某计划分商品实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-24&#34;,&#34;campaign_id_list&#34;:[2853805001],&#34;end_time&#34;:&#34;2021-09-24&#34;,&#34;launch_product_id_list&#34;:[101011001,101001005,101001013,101001014,101016001]}
type TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel is 查询某计划分商品实时报表 成功返回结果
type TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_material_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportMaterialRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse
func GetTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse() *TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse {
	return poolTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse.Get().(*TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse 将 TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse(v *TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse.Put(v)
}
