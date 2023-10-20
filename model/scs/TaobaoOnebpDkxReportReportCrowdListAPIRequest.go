package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCrowdListAPIRequest 获取人群离线报表 API请求
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;effect&#34;:15,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;crowd_id&#34;:12297883}
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxReportReportCrowdListAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._reportQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.crowd.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxReportReportCrowdListAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoOnebpDkxReportReportCrowdListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxReportReportCrowdListRequest()
	},
}

// GetTaobaoOnebpDkxReportReportCrowdListRequest 从 sync.Pool 获取 TaobaoOnebpDkxReportReportCrowdListAPIRequest
func GetTaobaoOnebpDkxReportReportCrowdListAPIRequest() *TaobaoOnebpDkxReportReportCrowdListAPIRequest {
	return poolTaobaoOnebpDkxReportReportCrowdListAPIRequest.Get().(*TaobaoOnebpDkxReportReportCrowdListAPIRequest)
}

// ReleaseTaobaoOnebpDkxReportReportCrowdListAPIRequest 将 TaobaoOnebpDkxReportReportCrowdListAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportCrowdListAPIRequest(v *TaobaoOnebpDkxReportReportCrowdListAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportCrowdListAPIRequest.Put(v)
}
