package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorFavoritesAPIRequest
手淘开放收藏夹鉴权接口 API请求
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。 */
type AlibabaInteractSensorFavoritesAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorFavoritesRequest 初始化AlibabaInteractSensorFavoritesAPIRequest对象
func NewAlibabaInteractSensorFavoritesRequest() *AlibabaInteractSensorFavoritesAPIRequest {
	return &AlibabaInteractSensorFavoritesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorFavoritesAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.favorites"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorFavoritesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
