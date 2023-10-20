package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest 查询某计划离线列表 API请求
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;launch_product_id_list&#34;:[101004013],&#34;start_time&#34;:&#34;2021-04-26&#34;,&#34;campaign_id_list&#34;:[134821085],&#34;end_time&#34;:&#34;2021-04-28&#34;,&#34;effect&#34;:15,}
// 非拓展流量查询：
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811599],&#34;end_time&#34;:&#34;2021-09-08&#34;,&#34;effect&#34;:15}
type TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCampaignOfflineRequest 初始化TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest对象
func NewTaobaoOnebpDkxReportReportCampaignOfflineRequest() *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest {
	return &TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
