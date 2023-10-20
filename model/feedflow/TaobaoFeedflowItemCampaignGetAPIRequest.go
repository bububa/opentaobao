package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignGetAPIRequest 通过计划id查询计划 API请求
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
type TaobaoFeedflowItemCampaignGetAPIRequest struct {
	model.Params
	// 计划id
	_campaginId int64
}

// NewTaobaoFeedflowItemCampaignGetRequest 初始化TaobaoFeedflowItemCampaignGetAPIRequest对象
func NewTaobaoFeedflowItemCampaignGetRequest() *TaobaoFeedflowItemCampaignGetAPIRequest {
	return &TaobaoFeedflowItemCampaignGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignGetAPIRequest) Reset() {
	r._campaginId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaginId is CampaginId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignGetAPIRequest) SetCampaginId(_campaginId int64) error {
	r._campaginId = _campaginId
	r.Set("campagin_id", _campaginId)
	return nil
}

// GetCampaginId CampaginId Getter
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetCampaginId() int64 {
	return r._campaginId
}

var poolTaobaoFeedflowItemCampaignGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignGetRequest()
	},
}

// GetTaobaoFeedflowItemCampaignGetRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignGetAPIRequest
func GetTaobaoFeedflowItemCampaignGetAPIRequest() *TaobaoFeedflowItemCampaignGetAPIRequest {
	return poolTaobaoFeedflowItemCampaignGetAPIRequest.Get().(*TaobaoFeedflowItemCampaignGetAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignGetAPIRequest 将 TaobaoFeedflowItemCampaignGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignGetAPIRequest(v *TaobaoFeedflowItemCampaignGetAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignGetAPIRequest.Put(v)
}
