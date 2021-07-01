package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单元id删除单元 API请求
taobao.feedflow.item.adgroup.delete

根据单元id删除单元
*/
type TaobaoFeedflowItemAdgroupDeleteAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 单元id列表
    _adgroupIdList   []int64
}

// 初始化TaobaoFeedflowItemAdgroupDeleteAPIRequest对象
func NewTaobaoFeedflowItemAdgroupDeleteRequest() *TaobaoFeedflowItemAdgroupDeleteAPIRequest{
    return &TaobaoFeedflowItemAdgroupDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupIdList Setter
// 单元id列表
func (r *TaobaoFeedflowItemAdgroupDeleteAPIRequest) SetAdgroupIdList(_adgroupIdList []int64) error {
    r._adgroupIdList = _adgroupIdList
    r.Set("adgroup_id_list", _adgroupIdList)
    return nil
}

// AdgroupIdList Getter
func (r TaobaoFeedflowItemAdgroupDeleteAPIRequest) GetAdgroupIdList() []int64 {
    return r._adgroupIdList
}
