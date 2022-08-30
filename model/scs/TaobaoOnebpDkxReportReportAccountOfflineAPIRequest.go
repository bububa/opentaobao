package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportAccountOfflineAPIRequest 获取账户历史报表 API请求
// taobao.onebp.dkx.report.report.account.offline
//
// 获取账户历史报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{     "start_time": "2021-07-24",     "effect": 15,     "end_time": "2021-08-21",     "strategy_scene":true,     "unify_type":"kuan",     "bizCode":"adStrategyDkx"   }
type TaobaoOnebpDkxReportReportAccountOfflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportAccountOfflineRequest 初始化TaobaoOnebpDkxReportReportAccountOfflineAPIRequest对象
func NewTaobaoOnebpDkxReportReportAccountOfflineRequest() *TaobaoOnebpDkxReportReportAccountOfflineAPIRequest {
	return &TaobaoOnebpDkxReportReportAccountOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.account.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportAccountOfflineAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
