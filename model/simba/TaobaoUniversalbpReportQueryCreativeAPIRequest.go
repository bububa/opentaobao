package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerycreativeAPIRequest 创意报表查询 API请求
// taobao.universalbp.report.query.creative
//
// 创意报表查询
type TaobaouniversalbpreportquerycreativeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topCreativeReportQueryVO
	_topCreativeReportQueryVO *TopCreativeReportQueryVo
}

// NewTaobaouniversalbpreportquerycreativeRequest 初始化TaobaouniversalbpreportquerycreativeAPIRequest对象
func NewTaobaouniversalbpreportquerycreativeRequest() *TaobaouniversalbpreportquerycreativeAPIRequest {
	return &TaobaouniversalbpreportquerycreativeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportquerycreativeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.creative"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportquerycreativeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportquerycreativeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportquerycreativeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportquerycreativeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopCreativeReportQueryVO is TopCreativeReportQueryVO Setter
// topCreativeReportQueryVO
func (r *TaobaouniversalbpreportquerycreativeAPIRequest) SetTopCreativeReportQueryVO(_topCreativeReportQueryVO *TopCreativeReportQueryVo) error {
	r._topCreativeReportQueryVO = _topCreativeReportQueryVO
	r.Set("top_creative_report_query_v_o", _topCreativeReportQueryVO)
	return nil
}

// GetTopCreativeReportQueryVO TopCreativeReportQueryVO Getter
func (r TaobaouniversalbpreportquerycreativeAPIRequest) GetTopCreativeReportQueryVO() *TopCreativeReportQueryVo {
	return r._topCreativeReportQueryVO
}
