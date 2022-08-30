package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListAPIResponse 获取人群离线报表 API返回值
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811613],"effect":15,"end_time":"2021-09-10","crowd_id":12297883}
type TaobaoOnebpDkxReportReportCrowdListAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCrowdListAPIResponseModel
}

// TaobaoOnebpDkxReportReportCrowdListAPIResponseModel is 获取人群离线报表 成功返回结果
type TaobaoOnebpDkxReportReportCrowdListAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_crowd_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCrowdListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
