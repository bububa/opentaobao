package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryaccountAPIRequest 账户报表查询 API请求
// taobao.universalbp.report.query.account
//
// 账户报表查询
type TaobaouniversalbpreportqueryaccountAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAccountReportQueryVO
	_topAccountReportQueryVO *TopAccountReportQueryVo
}

// NewTaobaouniversalbpreportqueryaccountRequest 初始化TaobaouniversalbpreportqueryaccountAPIRequest对象
func NewTaobaouniversalbpreportqueryaccountRequest() *TaobaouniversalbpreportqueryaccountAPIRequest {
	return &TaobaouniversalbpreportqueryaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportqueryaccountAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportqueryaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportqueryaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportqueryaccountAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportqueryaccountAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAccountReportQueryVO is TopAccountReportQueryVO Setter
// topAccountReportQueryVO
func (r *TaobaouniversalbpreportqueryaccountAPIRequest) SetTopAccountReportQueryVO(_topAccountReportQueryVO *TopAccountReportQueryVo) error {
	r._topAccountReportQueryVO = _topAccountReportQueryVO
	r.Set("top_account_report_query_v_o", _topAccountReportQueryVO)
	return nil
}

// GetTopAccountReportQueryVO TopAccountReportQueryVO Getter
func (r TaobaouniversalbpreportqueryaccountAPIRequest) GetTopAccountReportQueryVO() *TopAccountReportQueryVo {
	return r._topAccountReportQueryVO
}
