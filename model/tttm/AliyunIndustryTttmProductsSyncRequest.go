package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖货品信息同步 API请求
aliyun.industry.tttm.products.sync

天天特卖货品信息同步
*/
type AliyunIndustryTttmProductsSyncRequest struct {
    model.Params
    // 产品信息
    _syncProducts   []ProductInfoDTO
}

// 初始化AliyunIndustryTttmProductsSyncRequest对象
func NewAliyunIndustryTttmProductsSyncRequest() *AliyunIndustryTttmProductsSyncRequest{
    return &AliyunIndustryTttmProductsSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmProductsSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.products.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmProductsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncProducts Setter
// 产品信息
func (r *AliyunIndustryTttmProductsSyncRequest) SetSyncProducts(_syncProducts []ProductInfoDTO) error {
    r._syncProducts = _syncProducts
    r.Set("sync_products", _syncProducts)
    return nil
}

// SyncProducts Getter
func (r AliyunIndustryTttmProductsSyncRequest) GetSyncProducts() []ProductInfoDTO {
    return r._syncProducts
}
