package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportaccountofflineAPIRequest 获取账户历史报表 API请求
// taobao.onebp.dkx.report.report.account.offline
//
// 获取账户历史报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;start_time&#34;: &#34;2021-07-24&#34;,     &#34;effect&#34;: 15,     &#34;end_time&#34;: &#34;2021-08-21&#34;,     &#34;strategy_scene&#34;:true,     &#34;unify_type&#34;:&#34;kuan&#34;,     &#34;bizCode&#34;:&#34;adStrategyDkx&#34;   }
type TaobaoonebpdkxreportreportaccountofflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoonebpdkxreportreportaccountofflineRequest 初始化TaobaoonebpdkxreportreportaccountofflineAPIRequest对象
func NewTaobaoonebpdkxreportreportaccountofflineRequest() *TaobaoonebpdkxreportreportaccountofflineAPIRequest {
	return &TaobaoonebpdkxreportreportaccountofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxreportreportaccountofflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.account.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxreportreportaccountofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxreportreportaccountofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxreportreportaccountofflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxreportreportaccountofflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoonebpdkxreportreportaccountofflineAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoonebpdkxreportreportaccountofflineAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
