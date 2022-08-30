package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest 查询某计划实时列表 API请求
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"log_date_list": ["2021-09-09"],     "campaign_id_list": [2821811599]}
type TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCampaignRealtimeRequest 初始化TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest对象
func NewTaobaoOnebpDkxReportReportCampaignRealtimeRequest() *TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest {
	return &TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
