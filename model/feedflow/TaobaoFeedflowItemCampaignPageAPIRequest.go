package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignPageAPIRequest 批量查询计划列表 API请求
// taobao.feedflow.item.campaign.page
//
// 批量查询计划列表
type TaobaoFeedflowItemCampaignPageAPIRequest struct {
	model.Params
	// 入参
	_campaignQuery *CampaignQueryDto
}

// NewTaobaoFeedflowItemCampaignPageRequest 初始化TaobaoFeedflowItemCampaignPageAPIRequest对象
func NewTaobaoFeedflowItemCampaignPageRequest() *TaobaoFeedflowItemCampaignPageAPIRequest {
	return &TaobaoFeedflowItemCampaignPageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignPageAPIRequest) Reset() {
	r._campaignQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignQuery is CampaignQuery Setter
// 入参
func (r *TaobaoFeedflowItemCampaignPageAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r TaobaoFeedflowItemCampaignPageAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}

var poolTaobaoFeedflowItemCampaignPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignPageRequest()
	},
}

// GetTaobaoFeedflowItemCampaignPageRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignPageAPIRequest
func GetTaobaoFeedflowItemCampaignPageAPIRequest() *TaobaoFeedflowItemCampaignPageAPIRequest {
	return poolTaobaoFeedflowItemCampaignPageAPIRequest.Get().(*TaobaoFeedflowItemCampaignPageAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignPageAPIRequest 将 TaobaoFeedflowItemCampaignPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignPageAPIRequest(v *TaobaoFeedflowItemCampaignPageAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignPageAPIRequest.Put(v)
}
