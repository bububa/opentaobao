package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignRptdailylistAPIRequest 推广计划分日数据查询 API请求
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
type TaobaoFeedflowItemCampaignRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemCampaignRptdailylistRequest 初始化TaobaoFeedflowItemCampaignRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCampaignRptdailylistRequest() *TaobaoFeedflowItemCampaignRptdailylistAPIRequest {
	return &TaobaoFeedflowItemCampaignRptdailylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignRptdailylistAPIRequest) Reset() {
	r._rptQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCampaignRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}

var poolTaobaoFeedflowItemCampaignRptdailylistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignRptdailylistRequest()
	},
}

// GetTaobaoFeedflowItemCampaignRptdailylistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignRptdailylistAPIRequest
func GetTaobaoFeedflowItemCampaignRptdailylistAPIRequest() *TaobaoFeedflowItemCampaignRptdailylistAPIRequest {
	return poolTaobaoFeedflowItemCampaignRptdailylistAPIRequest.Get().(*TaobaoFeedflowItemCampaignRptdailylistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignRptdailylistAPIRequest 将 TaobaoFeedflowItemCampaignRptdailylistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignRptdailylistAPIRequest(v *TaobaoFeedflowItemCampaignRptdailylistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignRptdailylistAPIRequest.Put(v)
}
