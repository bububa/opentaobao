package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListAPIRequest 获取人群离线报表 API请求
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811613],"effect":15,"end_time":"2021-09-10","crowd_id":12297883}
type TaobaoOnebpDkxReportReportCrowdListAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCrowdListRequest 初始化TaobaoOnebpDkxReportReportCrowdListAPIRequest对象
func NewTaobaoOnebpDkxReportReportCrowdListRequest() *TaobaoOnebpDkxReportReportCrowdListAPIRequest {
	return &TaobaoOnebpDkxReportReportCrowdListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.crowd.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCrowdListAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCrowdListAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
