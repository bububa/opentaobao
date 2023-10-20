package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerycrowdAPIRequest 人群报表查询 API请求
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
type TaobaouniversalbpreportquerycrowdAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topCrowdReportQueryVO
	_topCrowdReportQueryVO *TopCrowdReportQueryVo
}

// NewTaobaouniversalbpreportquerycrowdRequest 初始化TaobaouniversalbpreportquerycrowdAPIRequest对象
func NewTaobaouniversalbpreportquerycrowdRequest() *TaobaouniversalbpreportquerycrowdAPIRequest {
	return &TaobaouniversalbpreportquerycrowdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportquerycrowdAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.crowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportquerycrowdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportquerycrowdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportquerycrowdAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportquerycrowdAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopCrowdReportQueryVO is TopCrowdReportQueryVO Setter
// topCrowdReportQueryVO
func (r *TaobaouniversalbpreportquerycrowdAPIRequest) SetTopCrowdReportQueryVO(_topCrowdReportQueryVO *TopCrowdReportQueryVo) error {
	r._topCrowdReportQueryVO = _topCrowdReportQueryVO
	r.Set("top_crowd_report_query_v_o", _topCrowdReportQueryVO)
	return nil
}

// GetTopCrowdReportQueryVO TopCrowdReportQueryVO Getter
func (r TaobaouniversalbpreportquerycrowdAPIRequest) GetTopCrowdReportQueryVO() *TopCrowdReportQueryVo {
	return r._topCrowdReportQueryVO
}
