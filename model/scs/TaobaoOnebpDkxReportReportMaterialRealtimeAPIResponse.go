package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse 查询某计划分商品实时报表 API返回值
// taobao.onebp.dkx.report.report.material.realtime
//
// 查询某计划分商品实时报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-24","campaign_id_list":[2853805001],"end_time":"2021-09-24","launch_product_id_list":[101011001,101001005,101001013,101001014,101016001]}
type TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel
}

// TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel is 查询某计划分商品实时报表 成功返回结果
type TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_material_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportMaterialRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
