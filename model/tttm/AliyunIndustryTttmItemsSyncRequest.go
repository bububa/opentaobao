package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖商品信息同步 API请求
aliyun.industry.tttm.items.sync

天天特卖商品信息同步
*/
type AliyunIndustryTttmItemsSyncAPIRequest struct {
    model.Params
    // 商品信息
    _syncItems   []ItemInfoDTO
}

// 初始化AliyunIndustryTttmItemsSyncAPIRequest对象
func NewAliyunIndustryTttmItemsSyncRequest() *AliyunIndustryTttmItemsSyncAPIRequest{
    return &AliyunIndustryTttmItemsSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.items.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncItems Setter
// 商品信息
func (r *AliyunIndustryTttmItemsSyncAPIRequest) SetSyncItems(_syncItems []ItemInfoDTO) error {
    r._syncItems = _syncItems
    r.Set("sync_items", _syncItems)
    return nil
}

// SyncItems Getter
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetSyncItems() []ItemInfoDTO {
    return r._syncItems
}
