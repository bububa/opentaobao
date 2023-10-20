package tttm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmItemsSyncAPIRequest 天天特卖商品信息同步 API请求
// aliyun.industry.tttm.items.sync
//
// 天天特卖商品信息同步
type AliyunIndustryTttmItemsSyncAPIRequest struct {
	model.Params
	// 商品信息
	_syncItems []ItemInfoDto
}

// NewAliyunIndustryTttmItemsSyncRequest 初始化AliyunIndustryTttmItemsSyncAPIRequest对象
func NewAliyunIndustryTttmItemsSyncRequest() *AliyunIndustryTttmItemsSyncAPIRequest {
	return &AliyunIndustryTttmItemsSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunIndustryTttmItemsSyncAPIRequest) Reset() {
	r._syncItems = r._syncItems[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.items.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncItems is SyncItems Setter
// 商品信息
func (r *AliyunIndustryTttmItemsSyncAPIRequest) SetSyncItems(_syncItems []ItemInfoDto) error {
	r._syncItems = _syncItems
	r.Set("sync_items", _syncItems)
	return nil
}

// GetSyncItems SyncItems Getter
func (r AliyunIndustryTttmItemsSyncAPIRequest) GetSyncItems() []ItemInfoDto {
	return r._syncItems
}

var poolAliyunIndustryTttmItemsSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunIndustryTttmItemsSyncRequest()
	},
}

// GetAliyunIndustryTttmItemsSyncRequest 从 sync.Pool 获取 AliyunIndustryTttmItemsSyncAPIRequest
func GetAliyunIndustryTttmItemsSyncAPIRequest() *AliyunIndustryTttmItemsSyncAPIRequest {
	return poolAliyunIndustryTttmItemsSyncAPIRequest.Get().(*AliyunIndustryTttmItemsSyncAPIRequest)
}

// ReleaseAliyunIndustryTttmItemsSyncAPIRequest 将 AliyunIndustryTttmItemsSyncAPIRequest 放入 sync.Pool
func ReleaseAliyunIndustryTttmItemsSyncAPIRequest(v *AliyunIndustryTttmItemsSyncAPIRequest) {
	v.Reset()
	poolAliyunIndustryTttmItemsSyncAPIRequest.Put(v)
}
