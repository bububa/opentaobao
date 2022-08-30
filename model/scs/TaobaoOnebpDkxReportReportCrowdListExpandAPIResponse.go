package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListExpandAPIResponse 获取拓展人群数据报表 API返回值
// taobao.onebp.dkx.report.report.crowd.list.expand
//
// 获取拓展人群数据报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"effect":15,"start_time":"2021-09-08","end_time":"2021-09-10","campaign_id_list":[2821811613],"white_crowd_id_List":[12297883,12298696,12297989]}
type TaobaoOnebpDkxReportReportCrowdListExpandAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCrowdListExpandAPIResponseModel
}

// TaobaoOnebpDkxReportReportCrowdListExpandAPIResponseModel is 获取拓展人群数据报表 成功返回结果
type TaobaoOnebpDkxReportReportCrowdListExpandAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_crowd_list_expand_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCrowdListExpandResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
