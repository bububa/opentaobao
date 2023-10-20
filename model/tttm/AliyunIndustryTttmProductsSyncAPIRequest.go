package tttm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmProductsSyncAPIRequest 天天特卖货品信息同步 API请求
// aliyun.industry.tttm.products.sync
//
// 天天特卖货品信息同步
type AliyunIndustryTttmProductsSyncAPIRequest struct {
	model.Params
	// 产品信息
	_syncProducts []ProductInfoDto
}

// NewAliyunIndustryTttmProductsSyncRequest 初始化AliyunIndustryTttmProductsSyncAPIRequest对象
func NewAliyunIndustryTttmProductsSyncRequest() *AliyunIndustryTttmProductsSyncAPIRequest {
	return &AliyunIndustryTttmProductsSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunIndustryTttmProductsSyncAPIRequest) Reset() {
	r._syncProducts = r._syncProducts[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmProductsSyncAPIRequest) GetApiMethodName() string {
	return "aliyun.industry.tttm.products.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmProductsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunIndustryTttmProductsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProducts is SyncProducts Setter
// 产品信息
func (r *AliyunIndustryTttmProductsSyncAPIRequest) SetSyncProducts(_syncProducts []ProductInfoDto) error {
	r._syncProducts = _syncProducts
	r.Set("sync_products", _syncProducts)
	return nil
}

// GetSyncProducts SyncProducts Getter
func (r AliyunIndustryTttmProductsSyncAPIRequest) GetSyncProducts() []ProductInfoDto {
	return r._syncProducts
}

var poolAliyunIndustryTttmProductsSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunIndustryTttmProductsSyncRequest()
	},
}

// GetAliyunIndustryTttmProductsSyncRequest 从 sync.Pool 获取 AliyunIndustryTttmProductsSyncAPIRequest
func GetAliyunIndustryTttmProductsSyncAPIRequest() *AliyunIndustryTttmProductsSyncAPIRequest {
	return poolAliyunIndustryTttmProductsSyncAPIRequest.Get().(*AliyunIndustryTttmProductsSyncAPIRequest)
}

// ReleaseAliyunIndustryTttmProductsSyncAPIRequest 将 AliyunIndustryTttmProductsSyncAPIRequest 放入 sync.Pool
func ReleaseAliyunIndustryTttmProductsSyncAPIRequest(v *AliyunIndustryTttmProductsSyncAPIRequest) {
	v.Reset()
	poolAliyunIndustryTttmProductsSyncAPIRequest.Put(v)
}
