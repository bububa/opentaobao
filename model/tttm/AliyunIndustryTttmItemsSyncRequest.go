package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖商品信息同步 APIRequest
aliyun.industry.tttm.items.sync

天天特卖商品信息同步
*/
type AliyunIndustryTttmItemsSyncRequest struct {
    model.Params

    // 商品信息
    syncItems   []ItemInfoDto 

}

func NewAliyunIndustryTttmItemsSyncRequest() *AliyunIndustryTttmItemsSyncRequest{
    return &AliyunIndustryTttmItemsSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmItemsSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.items.sync"
}

func (r AliyunIndustryTttmItemsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmItemsSyncRequest) SetSyncItems(syncItems []ItemInfoDto) error {
    r.syncItems = syncItems
    r.Set("sync_items", syncItems)
    return nil
}

func (r AliyunIndustryTttmItemsSyncRequest) GetSyncItems() []ItemInfoDto {
    return r.syncItems
}

