package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryBidwordAPIRequest 关键词报表查询 API请求
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
type TaobaoUniversalbpReportQueryBidwordAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topBidWordReportQueryVO
	_topBidWordReportQueryVO *TopBidWordReportQueryVo
}

// NewTaobaoUniversalbpReportQueryBidwordRequest 初始化TaobaoUniversalbpReportQueryBidwordAPIRequest对象
func NewTaobaoUniversalbpReportQueryBidwordRequest() *TaobaoUniversalbpReportQueryBidwordAPIRequest {
	return &TaobaoUniversalbpReportQueryBidwordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryBidwordAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.bidword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryBidwordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryBidwordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryBidwordAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryBidwordAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopBidWordReportQueryVO is TopBidWordReportQueryVO Setter
// topBidWordReportQueryVO
func (r *TaobaoUniversalbpReportQueryBidwordAPIRequest) SetTopBidWordReportQueryVO(_topBidWordReportQueryVO *TopBidWordReportQueryVo) error {
	r._topBidWordReportQueryVO = _topBidWordReportQueryVO
	r.Set("top_bid_word_report_query_v_o", _topBidWordReportQueryVO)
	return nil
}

// GetTopBidWordReportQueryVO TopBidWordReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryBidwordAPIRequest) GetTopBidWordReportQueryVO() *TopBidWordReportQueryVo {
	return r._topBidWordReportQueryVO
}
