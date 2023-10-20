package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCreativeAPIRequest 创意报表查询 API请求
// taobao.universalbp.report.query.creative
//
// 创意报表查询
type TaobaoUniversalbpReportQueryCreativeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topCreativeReportQueryVO
	_topCreativeReportQueryVO *TopCreativeReportQueryVo
}

// NewTaobaoUniversalbpReportQueryCreativeRequest 初始化TaobaoUniversalbpReportQueryCreativeAPIRequest对象
func NewTaobaoUniversalbpReportQueryCreativeRequest() *TaobaoUniversalbpReportQueryCreativeAPIRequest {
	return &TaobaoUniversalbpReportQueryCreativeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryCreativeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.creative"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryCreativeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryCreativeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryCreativeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryCreativeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopCreativeReportQueryVO is TopCreativeReportQueryVO Setter
// topCreativeReportQueryVO
func (r *TaobaoUniversalbpReportQueryCreativeAPIRequest) SetTopCreativeReportQueryVO(_topCreativeReportQueryVO *TopCreativeReportQueryVo) error {
	r._topCreativeReportQueryVO = _topCreativeReportQueryVO
	r.Set("top_creative_report_query_v_o", _topCreativeReportQueryVO)
	return nil
}

// GetTopCreativeReportQueryVO TopCreativeReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryCreativeAPIRequest) GetTopCreativeReportQueryVO() *TopCreativeReportQueryVo {
	return r._topCreativeReportQueryVO
}
