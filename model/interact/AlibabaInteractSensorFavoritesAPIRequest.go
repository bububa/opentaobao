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

// New
