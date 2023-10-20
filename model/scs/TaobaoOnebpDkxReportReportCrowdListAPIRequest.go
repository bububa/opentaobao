package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcrowdlistAPIRequest 获取人群离线报表 API请求
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;effect&#34;:15,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;crowd_id&#34;:12297883}
type TaobaoonebpdkxreportreportcrowdlistAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoonebpdkxreportreportcrowdlistRequest 初始化TaobaoonebpdkxreportreportcrowdlistAPIRequest对象
func NewTaobaoonebpdkxreportreportcrowdlistRequest() *TaobaoonebpdkxreportreportcrowdlistAPIRequest {
	return &TaobaoonebpdkxreportreportcrowdlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxreportreportcrowdlistAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.crowd.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxreportreportcrowdlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxreportreportcrowdlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxreportreportcrowdlistAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxreportreportcrowdlistAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoonebpdkxreportreportcrowdlistAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoonebpdkxreportreportcrowdlistAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
