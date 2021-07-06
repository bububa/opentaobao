package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupDeleteAPIRequest 根据单元id删除单元 API请求
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
type TaobaoFeedflowItemAdgroupDeleteAPIRequest struct {
	model.Params
	// 单元id列表
	_adgroupIdList []int64
	// 计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemAdgroupDeleteRequest 初始化TaobaoFeedflowItemAdgroupDeleteAPIRequest对象
func NewTaobaoFeedflowItemAdgroupDeleteRequest() *TaobaoFeedflowItemAdgroupDeleteAPIRequest {
	return &TaobaoFeedflowItemAdgroupDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdgroupIdList is AdgroupIdList Setter
// 单元id列表
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) SetAdgroupIdList(_adgroupIdList []int64) error {
	r._adgroupIdList = _adgroupIdList
	r.Set("adgroup_id_list", _adgroupIdList)
	return nil
}

// GetAdgroupIdList AdgroupIdList Getter
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetAdgroupIdList() []int64 {
	return r._adgroupIdList
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
