package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmstocksyncAPIRequest 天天特卖库存同步接口 API请求
// aliyun.industry.tttm.stock.sync
//
// 天天特卖库存同步接口
type AliyunindustrytttmstocksyncAPIRequest struct {
	model.Params
	// 库存
	_syncStock *StockInfoDto
}

// NewAliyunindustrytttmstocksyncRequest 初始化AliyunindustrytttmstocksyncAPIRequest对象
func NewAliyunindustrytttmstocksyncRequest() *AliyunindustrytttmstocksyncAPIRequest {
	return &AliyunindustrytttmstocksyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunindustrytttmstocksyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.stock.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunindustrytttmstocksyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunindustrytttmstocksyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncStock is SyncStock Setter
// 库存
func (r *AliyunindustrytttmstocksyncAPIRequest) SetSyncStock(_syncStock *StockInfoDto) error {
	r._syncStock = _syncStock
	r.Set("sync_stock", _syncStock)
	return nil
}

// GetSyncStock SyncStock Getter
func (r AliyunindustrytttmstocksyncAPIRequest) GetSyncStock() *StockInfoDto {
	return r._syncStock
}
