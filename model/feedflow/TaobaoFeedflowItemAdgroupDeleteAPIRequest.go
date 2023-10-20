package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupDeleteAPIRequest 根据单元id删除单元 API请求
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
type TaobaoFeedflowItemAdgroupDeleteAPIRequest struct {
	model.Params
	// 单元id列表
	_adgroupIdList []string
	// 计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemAdgroupDeleteRequest 初始化TaobaoFeedflowItemAdgroupDeleteAPIRequest对象
func NewTaobaoFeedflowItemAdgroupDeleteRequest() *TaobaoFeedflowItemAdgroupDeleteAPIRequest {
	return &TaobaoFeedflowItemAdgroupDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) Reset() {
	r._adgroupIdList = r._adgroupIdList[:0]
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdList is AdgroupIdList Setter
// 单元id列表
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) SetAdgroupIdList(_adgroupIdList []string) error {
	r._adgroupIdList = _adgroupIdList
	r.Set("adgroup_id_list", _adgroupIdList)
	return nil
}

// GetAdgroupIdList AdgroupIdList Getter
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetAdgroupIdList() []string {
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

var poolTaobaoFeedflowItemAdgroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupDeleteRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupDeleteRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupDeleteAPIRequest
func GetTaobaoFeedflowItemAdgroupDeleteAPIRequest() *TaobaoFeedflowItemAdgroupDeleteAPIRequest {
	return poolTaobaoFeedflowItemAdgroupDeleteAPIRequest.Get().(*TaobaoFeedflowItemAdgroupDeleteAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupDeleteAPIRequest 将 TaobaoFeedflowItemAdgroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupDeleteAPIRequest(v *TaobaoFeedflowItemAdgroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupDeleteAPIRequest.Put(v)
}
