package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorFavorites 手淘开放收藏夹鉴权接口
// alibaba.interact.sensor.favorites
//
// 手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
func AlibabaInteractSensorFavorites(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorFavoritesAPIRequest, resp *interact.AlibabaInteractSensorFavoritesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
