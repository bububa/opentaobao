package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignRpthourlistAPIRequest 超级推荐【商品推广】计划分时报表查询 API请求
// taobao.feedflow.item.campaign.rpthourlist
//
// 广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据
type TaobaoFeedflowItemCampaignRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaoFeedflowItemCampaignRpthourlistRequest 初始化TaobaoFeedflowItemCampaignRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemCampaignRpthourlistRequest() *TaobaoFeedflowItemCampaignRpthourlistAPIRequest {
	return &TaobaoFeedflowItemCampaignRpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCampaignRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
