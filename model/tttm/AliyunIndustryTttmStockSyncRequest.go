package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖库存同步接口 APIRequest
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口
*/
type AliyunIndustryTttmStockSyncRequest struct {
    model.Params

    // 库存
    syncStock   *StockInfoDto 

}

func NewAliyunIndustryTttmStockSyncRequest() *AliyunIndustryTttmStockSyncRequest{
    return &AliyunIndustryTttmStockSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmStockSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.stock.sync"
}

func (r AliyunIndustryTttmStockSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmStockSyncRequest) SetSyncStock(syncStock *StockInfoDto) error {
    r.syncStock = syncStock
    r.Set("sync_stock", syncStock)
    return nil
}

func (r AliyunIndustryTttmStockSyncRequest) GetSyncStock() *StockInfoDto {
    return r.syncStock
}

