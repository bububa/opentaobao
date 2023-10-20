package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemTargetValidlistAPIRequest 获取有权限的定向列表 API请求
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
type TaobaoFeedflowItemTargetValidlistAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemTargetValidlistRequest 初始化TaobaoFeedflowItemTargetValidlistAPIRequest对象
func NewTaobaoFeedflowItemTargetValidlistRequest() *TaobaoFeedflowItemTargetValidlistAPIRequest {
	return &TaobaoFeedflowItemTargetValidlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemTargetValidlistAPIRequest) Reset() {
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.target.validlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemTargetValidlistAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoFeedflowItemTargetValidlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemTargetValidlistRequest()
	},
}

// GetTaobaoFeedflowItemTargetValidlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemTargetValidlistAPIRequest
func GetTaobaoFeedflowItemTargetValidlistAPIRequest() *TaobaoFeedflowItemTargetValidlistAPIRequest {
	return poolTaobaoFeedflowItemTargetValidlistAPIRequest.Get().(*TaobaoFeedflowItemTargetValidlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemTargetValidlistAPIRequest 将 TaobaoFeedflowItemTargetValidlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemTargetValidlistAPIRequest(v *TaobaoFeedflowItemTargetValidlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemTargetValidlistAPIRequest.Put(v)
}
