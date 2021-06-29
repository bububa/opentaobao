package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖库存同步接口 API请求
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口
*/
type AliyunIndustryTttmStockSyncRequest struct {
    model.Params
    // 库存
    _syncStock   *StockInfoDTO
}

// 初始化AliyunIndustryTttmStockSyncRequest对象
func NewAliyunIndustryTttmStockSyncRequest() *AliyunIndustryTttmStockSyncRequest{
    return &AliyunIndustryTttmStockSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmStockSyncRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.stock.sync"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmStockSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncStock Setter
// 库存
func (r *AliyunIndustryTttmStockSyncRequest) SetSyncStock(_syncStock *StockInfoDTO) error {
    r._syncStock = _syncStock
    r.Set("sync_stock", _syncStock)
    return nil
}

// SyncStock Getter
func (r AliyunIndustryTttmStockSyncRequest) GetSyncStock() *StockInfoDTO {
    return r._syncStock
}
