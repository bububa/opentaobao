package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorfavoritesAPIRequest 手淘开放收藏夹鉴权接口 API请求
// alibaba.interact.sensor.favorites
//
// 手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
type AlibabainteractsensorfavoritesAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorfavoritesRequest 初始化AlibabainteractsensorfavoritesAPIRequest对象
func NewAlibabainteractsensorfavoritesRequest() *AlibabainteractsensorfavoritesAPIRequest {
	return &AlibabainteractsensorfavoritesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorfavoritesAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.favorites"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorfavoritesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorfavoritesAPIRequest) GetRawParams() model.Params {
	return r.Params
}
