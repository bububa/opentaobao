package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryareaAPIRequest 地域报表查询 API请求
// taobao.universalbp.report.query.area
//
// 地域报表查询
type TaobaouniversalbpreportqueryareaAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAreaReportQueryVO
	_topAreaReportQueryVO *TopAreaReportQueryVo
}

// NewTaobaouniversalbpreportqueryareaRequest 初始化TaobaouniversalbpreportqueryareaAPIRequest对象
func NewTaobaouniversalbpreportqueryareaRequest() *TaobaouniversalbpreportqueryareaAPIRequest {
	return &TaobaouniversalbpreportqueryareaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportqueryareaAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.area"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportqueryareaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportqueryareaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportqueryareaAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportqueryareaAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAreaReportQueryVO is TopAreaReportQueryVO Setter
// topAreaReportQueryVO
func (r *TaobaouniversalbpreportqueryareaAPIRequest) SetTopAreaReportQueryVO(_topAreaReportQueryVO *TopAreaReportQueryVo) error {
	r._topAreaReportQueryVO = _topAreaReportQueryVO
	r.Set("top_area_report_query_v_o", _topAreaReportQueryVO)
	return nil
}

// GetTopAreaReportQueryVO TopAreaReportQueryVO Getter
func (r TaobaouniversalbpreportqueryareaAPIRequest) GetTopAreaReportQueryVO() *TopAreaReportQueryVo {
	return r._topAreaReportQueryVO
}
