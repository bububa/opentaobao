package tttm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
