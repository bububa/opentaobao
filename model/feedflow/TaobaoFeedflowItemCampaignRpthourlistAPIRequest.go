package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignRpthourlistAPIRequest) Reset() {
	r._rptQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignRpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoFeedflowItemCampaignRpthourlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignRpthourlistRequest()
	},
}

// GetTaobaoFeedflowItemCampaignRpthourlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignRpthourlistAPIRequest
func GetTaobaoFeedflowItemCampaignRpthourlistAPIRequest() *TaobaoFeedflowItemCampaignRpthourlistAPIRequest {
	return poolTaobaoFeedflowItemCampaignRpthourlistAPIRequest.Get().(*TaobaoFeedflowItemCampaignRpthourlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignRpthourlistAPIRequest 将 TaobaoFeedflowItemCampaignRpthourlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignRpthourlistAPIRequest(v *TaobaoFeedflowItemCampaignRpthourlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignRpthourlistAPIRequest.Put(v)
}
