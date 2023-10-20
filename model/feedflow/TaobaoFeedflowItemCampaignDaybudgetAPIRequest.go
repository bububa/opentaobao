package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignDaybudgetAPIRequest 获取当日投放日预算总额 API请求
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
type TaobaoFeedflowItemCampaignDaybudgetAPIRequest struct {
	model.Params
}

// NewTaobaoFeedflowItemCampaignDaybudgetRequest 初始化TaobaoFeedflowItemCampaignDaybudgetAPIRequest对象
func NewTaobaoFeedflowItemCampaignDaybudgetRequest() *TaobaoFeedflowItemCampaignDaybudgetAPIRequest {
	return &TaobaoFeedflowItemCampaignDaybudgetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignDaybudgetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignDaybudgetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.daybudget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignDaybudgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignDaybudgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoFeedflowItemCampaignDaybudgetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignDaybudgetRequest()
	},
}

// GetTaobaoFeedflowItemCampaignDaybudgetRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignDaybudgetAPIRequest
func GetTaobaoFeedflowItemCampaignDaybudgetAPIRequest() *TaobaoFeedflowItemCampaignDaybudgetAPIRequest {
	return poolTaobaoFeedflowItemCampaignDaybudgetAPIRequest.Get().(*TaobaoFeedflowItemCampaignDaybudgetAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignDaybudgetAPIRequest 将 TaobaoFeedflowItemCampaignDaybudgetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignDaybudgetAPIRequest(v *TaobaoFeedflowItemCampaignDaybudgetAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignDaybudgetAPIRequest.Put(v)
}
