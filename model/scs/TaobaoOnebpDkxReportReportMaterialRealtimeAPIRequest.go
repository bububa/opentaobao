package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest 查询某计划分商品实时报表 API请求
// taobao.onebp.dkx.report.report.material.realtime
//
// 查询某计划分商品实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-24&#34;,&#34;campaign_id_list&#34;:[2853805001],&#34;end_time&#34;:&#34;2021-09-24&#34;,&#34;launch_product_id_list&#34;:[101011001,101001005,101001013,101001014,101016001]}
type TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportMaterialRealtimeRequest 初始化TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest对象
func NewTaobaoOnebpDkxReportReportMaterialRealtimeRequest() *TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest {
	return &TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.material.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
