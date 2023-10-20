package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryadgroupAPIRequest 单元报表查询 API请求
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
type TaobaouniversalbpreportqueryadgroupAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAdgroupReportQueryVO
	_topAdgroupReportQueryVO *TopAdgroupReportQueryVo
}

// NewTaobaouniversalbpreportqueryadgroupRequest 初始化TaobaouniversalbpreportqueryadgroupAPIRequest对象
func NewTaobaouniversalbpreportqueryadgroupRequest() *TaobaouniversalbpreportqueryadgroupAPIRequest {
	return &TaobaouniversalbpreportqueryadgroupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportqueryadgroupAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.adgroup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportqueryadgroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportqueryadgroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportqueryadgroupAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportqueryadgroupAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAdgroupReportQueryVO is TopAdgroupReportQueryVO Setter
// topAdgroupReportQueryVO
func (r *TaobaouniversalbpreportqueryadgroupAPIRequest) SetTopAdgroupReportQueryVO(_topAdgroupReportQueryVO *TopAdgroupReportQueryVo) error {
	r._topAdgroupReportQueryVO = _topAdgroupReportQueryVO
	r.Set("top_adgroup_report_query_v_o", _topAdgroupReportQueryVO)
	return nil
}

// GetTopAdgroupReportQueryVO TopAdgroupReportQueryVO Getter
func (r TaobaouniversalbpreportqueryadgroupAPIRequest) GetTopAdgroupReportQueryVO() *TopAdgroupReportQueryVo {
	return r._topAdgroupReportQueryVO
}
