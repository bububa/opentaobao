package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignrpthourlistAPIRequest 超级推荐【商品推广】计划分时报表查询 API请求
// taobao.feedflow.item.campaign.rpthourlist
//
// 广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据
type TaobaofeedflowitemcampaignrpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowitemcampaignrpthourlistRequest 初始化TaobaofeedflowitemcampaignrpthourlistAPIRequest对象
func NewTaobaofeedflowitemcampaignrpthourlistRequest() *TaobaofeedflowitemcampaignrpthourlistAPIRequest {
	return &TaobaofeedflowitemcampaignrpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaignrpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaignrpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaignrpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowitemcampaignrpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowitemcampaignrpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
