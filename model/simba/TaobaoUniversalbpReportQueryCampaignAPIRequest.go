package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCampaignAPIRequest 计划报表查询 API请求
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
type TaobaoUniversalbpReportQueryCampaignAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topCampaignReportQueryVO
	_topCampaignReportQueryVO *TopCampaignReportQueryVo
}

// NewTaobaoUniversalbpReportQueryCampaignRequest 初始化TaobaoUniversalbpReportQueryCampaignAPIRequest对象
func NewTaobaoUniversalbpReportQueryCampaignRequest() *TaobaoUniversalbpReportQueryCampaignAPIRequest {
	return &TaobaoUniversalbpReportQueryCampaignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryCampaignAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.campaign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryCampaignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryCampaignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryCampaignAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryCampaignAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopCampaignReportQueryVO is TopCampaignReportQueryVO Setter
// topCampaignReportQueryVO
func (r *TaobaoUniversalbpReportQueryCampaignAPIRequest) SetTopCampaignReportQueryVO(_topCampaignReportQueryVO *TopCampaignReportQueryVo) error {
	r._topCampaignReportQueryVO = _topCampaignReportQueryVO
	r.Set("top_campaign_report_query_v_o", _topCampaignReportQueryVO)
	return nil
}

// GetTopCampaignReportQueryVO TopCampaignReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryCampaignAPIRequest) GetTopCampaignReportQueryVO() *TopCampaignReportQueryVo {
	return r._topCampaignReportQueryVO
}
