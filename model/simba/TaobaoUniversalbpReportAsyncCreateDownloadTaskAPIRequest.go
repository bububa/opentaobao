package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest 创建异步下载任务 API请求
// taobao.universalbp.report.async.create.download.task
//
// 入参报表查询信息，出参下载任务id
type TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// reportQueryVO
	_reportQueryVO *ReportQueryVo
}

// NewTaobaouniversalbpreportasynccreatedownloadtaskRequest 初始化TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest对象
func NewTaobaouniversalbpreportasynccreatedownloadtaskRequest() *TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest {
	return &TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.async.create.download.task"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetReportQueryVO is ReportQueryVO Setter
// reportQueryVO
func (r *TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) SetReportQueryVO(_reportQueryVO *ReportQueryVo) error {
	r._reportQueryVO = _reportQueryVO
	r.Set("report_query_v_o", _reportQueryVO)
	return nil
}

// GetReportQueryVO ReportQueryVO Getter
func (r TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest) GetReportQueryVO() *ReportQueryVo {
	return r._reportQueryVO
}
