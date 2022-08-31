package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest 获取账户实时报表 API请求
// taobao.onebp.dkx.report.report.account.realtime
//
// 获取账户实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;log_date_list&#34;: [  &#34;2021-09-23&#34;     ]   }
type TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportAccountRealtimeRequest 初始化TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest对象
func NewTaobaoOnebpDkxReportReportAccountRealtimeRequest() *TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest {
	return &TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.account.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
