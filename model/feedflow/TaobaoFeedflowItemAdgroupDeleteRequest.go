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
type TaobaoFeedflowItemAdgroupDeleteRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 单元id列表
    _adgroupIdList   []int64
}

// 初始化TaobaoFeedflowItemAdgroupDeleteRequest对象
func NewTaobaoFeedflowItemAdgroupDeleteRequest() *TaobaoFeedflowItemAdgroupDeleteRequest{
    return &TaobaoFeedflowItemAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemAdgroupDeleteRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupIdList Setter
// 单元id列表
func (r *TaobaoFeedflowItemAdgroupDeleteRequest) SetAdgroupIdList(_adgroupIdList []int64) error {
    r._adgroupIdList = _adgroupIdList
    r.Set("adgroup_id_list", _adgroupIdList)
    return nil
}

// AdgroupIdList Getter
func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetAdgroupIdList() []int64 {
    return r._adgroupIdList
}
