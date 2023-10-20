package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryrealtimeAPIRequest 实时报表查询 API请求
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
type TaobaouniversalbpreportqueryrealtimeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topRealTimeReportQueryVO
	_topRealTimeReportQueryVO *TopRealTimeReportQueryVo
}

// NewTaobaouniversalbpreportqueryrealtimeRequest 初始化TaobaouniversalbpreportqueryrealtimeAPIRequest对象
func NewTaobaouniversalbpreportqueryrealtimeRequest() *TaobaouniversalbpreportqueryrealtimeAPIRequest {
	return &TaobaouniversalbpreportqueryrealtimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportqueryrealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportqueryrealtimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportqueryrealtimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportqueryrealtimeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportqueryrealtimeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopRealTimeReportQueryVO is TopRealTimeReportQueryVO Setter
// topRealTimeReportQueryVO
func (r *TaobaouniversalbpreportqueryrealtimeAPIRequest) SetTopRealTimeReportQueryVO(_topRealTimeReportQueryVO *TopRealTimeReportQueryVo) error {
	r._topRealTimeReportQueryVO = _topRealTimeReportQueryVO
	r.Set("top_real_time_report_query_v_o", _topRealTimeReportQueryVO)
	return nil
}

// GetTopRealTimeReportQueryVO TopRealTimeReportQueryVO Getter
func (r TaobaouniversalbpreportqueryrealtimeAPIRequest) GetTopRealTimeReportQueryVO() *TopRealTimeReportQueryVo {
	return r._topRealTimeReportQueryVO
}
