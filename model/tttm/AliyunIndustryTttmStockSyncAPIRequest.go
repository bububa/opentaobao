package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmStockSyncAPIRequest
天天特卖库存同步接口 API请求
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口 */
type AliyunIndustryTttmStockSyncAPIRequest struct {
	model.Params
	// 库存
	_syncStock *StockInfoDto
}

// NewAliyunIndustryTttmStockSyncRequest 初始化AliyunIndustryTttmStockSyncAPIRequest对象
func NewAliyunIndustryTttmStockSyncRequest() *AliyunIndustryTttmStockSyncAPIRequest {
	return &AliyunIndustryTttmStockSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmStockSyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.stock.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmStockSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SyncStock Setter
// 库存
func (r *AliyunIndustryTttmStockSyncAPIRequest) SetSyncStock(_syncStock *StockInfoDto) error {
	r._syncStock = _syncStock
	r.Set("sync_stock", _syncStock)
	return nil
}

// Get SyncStock Getter
func (r AliyunIndustryTttmStockSyncAPIRequest) GetSyncStock() *StockInfoDto {
	return r._syncStock
}
