package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportMaterialOfflineAPIResponse 查询某计划分商品离线报表 API返回值
// taobao.onebp.dkx.report.report.material.offline
//
// 查询某计划分商品离线报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-23","campaign_id_list":[2853805001],"end_time":"2021-09-24","effect": 15   }
type TaobaoOnebpDkxReportReportMaterialOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportMaterialOfflineAPIResponseModel
}

// TaobaoOnebpDkxReportReportMaterialOfflineAPIResponseModel is 查询某计划分商品离线报表 成功返回结果
type TaobaoOnebpDkxReportReportMaterialOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_material_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportMaterialOfflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
