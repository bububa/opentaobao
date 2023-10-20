package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorFavoritesAPIRequest 手淘开放收藏夹鉴权接口 API请求
// alibaba.interact.sensor.favorites
//
// 手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
type AlibabaInteractSensorFavoritesAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorFavoritesRequest 初始化AlibabaInteractSensorFavoritesAPIRequest对象
func NewAlibabaInteractSensorFavoritesRequest() *AlibabaInteractSensorFavoritesAPIRequest {
	return &AlibabaInteractSensorFavoritesAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorFavoritesAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorFavoritesAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.favorites"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorFavoritesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorFavoritesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorFavoritesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorFavoritesRequest()
	},
}

// GetAlibabaInteractSensorFavoritesRequest 从 sync.Pool 获取 AlibabaInteractSensorFavoritesAPIRequest
func GetAlibabaInteractSensorFavoritesAPIRequest() *AlibabaInteractSensorFavoritesAPIRequest {
	return poolAlibabaInteractSensorFavoritesAPIRequest.Get().(*AlibabaInteractSensorFavoritesAPIRequest)
}

// ReleaseAlibabaInteractSensorFavoritesAPIRequest 将 AlibabaInteractSensorFavoritesAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorFavoritesAPIRequest(v *AlibabaInteractSensorFavoritesAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorFavoritesAPIRequest.Put(v)
}
