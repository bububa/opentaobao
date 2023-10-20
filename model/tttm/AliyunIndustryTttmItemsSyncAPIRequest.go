package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmitemssyncAPIRequest 天天特卖商品信息同步 API请求
// aliyun.industry.tttm.items.sync
//
// 天天特卖商品信息同步
type AliyunindustrytttmitemssyncAPIRequest struct {
	model.Params
	// 商品信息
	_syncItems []ItemInfoDto
}

// NewAliyunindustrytttmitemssyncRequest 初始化AliyunindustrytttmitemssyncAPIRequest对象
func NewAliyunindustrytttmitemssyncRequest() *AliyunindustrytttmitemssyncAPIRequest {
	return &AliyunindustrytttmitemssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunindustrytttmitemssyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.items.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunindustrytttmitemssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunindustrytttmitemssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncItems is SyncItems Setter
// 商品信息
func (r *AliyunindustrytttmitemssyncAPIRequest) SetSyncItems(_syncItems []ItemInfoDto) error {
	r._syncItems = _syncItems
	r.Set("sync_items", _syncItems)
	return nil
}

// GetSyncItems SyncItems Getter
func (r AliyunindustrytttmitemssyncAPIRequest) GetSyncItems() []ItemInfoDto {
	return r._syncItems
}
