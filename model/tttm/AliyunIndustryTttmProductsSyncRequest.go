package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖货品信息同步 APIRequest
aliyun.industry.tttm.products.sync

天天特卖货品信息同步
*/
type AliyunIndustryTttmProductsSyncRequest struct {
    model.Params

    // 产品信息
    syncProducts   []ProductInfoDto 

}

func NewAliyunIndustryTttmProductsSyncRequest() *AliyunIndustryTttmProductsSyncRequest{
    return &AliyunIndustryTttmProductsSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmProductsSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.products.sync"
}

func (r AliyunIndustryTttmProductsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmProductsSyncRequest) SetSyncProducts(syncProducts []ProductInfoDto) error {
    r.syncProducts = syncProducts
    r.Set("sync_products", syncProducts)
    return nil
}

func (r AliyunIndustryTttmProductsSyncRequest) GetSyncProducts() []ProductInfoDto {
    return r.syncProducts
}

