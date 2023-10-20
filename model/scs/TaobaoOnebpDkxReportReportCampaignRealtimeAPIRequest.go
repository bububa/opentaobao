package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest 查询某计划实时列表 API请求
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;log_date_list&#34;: [&#34;2021-09-09&#34;],     &#34;campaign_id_list&#34;: [2821811599]}
type TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoonebpdkxreportreportcampaignrealtimeRequest 初始化TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest对象
func NewTaobaoonebpdkxreportreportcampaignrealtimeRequest() *TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest {
	return &TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoonebpdkxreportreportcampaignrealtimeAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
