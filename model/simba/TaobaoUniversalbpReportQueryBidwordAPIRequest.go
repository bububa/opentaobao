package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerybidwordAPIRequest 关键词报表查询 API请求
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
type TaobaouniversalbpreportquerybidwordAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topBidWordReportQueryVO
	_topBidWordReportQueryVO *TopBidWordReportQueryVo
}

// NewTaobaouniversalbpreportquerybidwordRequest 初始化TaobaouniversalbpreportquerybidwordAPIRequest对象
func NewTaobaouniversalbpreportquerybidwordRequest() *TaobaouniversalbpreportquerybidwordAPIRequest {
	return &TaobaouniversalbpreportquerybidwordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportquerybidwordAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.bidword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportquerybidwordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportquerybidwordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportquerybidwordAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportquerybidwordAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopBidWordReportQueryVO is TopBidWordReportQueryVO Setter
// topBidWordReportQueryVO
func (r *TaobaouniversalbpreportquerybidwordAPIRequest) SetTopBidWordReportQueryVO(_topBidWordReportQueryVO *TopBidWordReportQueryVo) error {
	r._topBidWordReportQueryVO = _topBidWordReportQueryVO
	r.Set("top_bid_word_report_query_v_o", _topBidWordReportQueryVO)
	return nil
}

// GetTopBidWordReportQueryVO TopBidWordReportQueryVO Getter
func (r TaobaouniversalbpreportquerybidwordAPIRequest) GetTopBidWordReportQueryVO() *TopBidWordReportQueryVo {
	return r._topBidWordReportQueryVO
}
