package feedflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaginId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignGetAPIRequest) SetCampaginId(_campaginId int64) error {
	r._campaginId = _campaginId
	r.Set("campagin_id", _campaginId)
	return nil
}

// Get CampaginId Getter
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetCampaginId() int64 {
	return r._campaginId
}
