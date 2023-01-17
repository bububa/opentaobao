package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest 获取拓展人群数据报表 API请求
// taobao.onebp.dkx.report.report.crowd.list.expand
//
// 获取拓展人群数据报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;effect&#34;:15,&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;white_crowd_id_List&#34;:[12297883,12298696,12297989]}
type TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCrowdListExpandRequest 初始化TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest对象
func NewTaobaoOnebpDkxReportReportCrowdListExpandRequest() *TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest {
	return &TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.crowd.list.expand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
