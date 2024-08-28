package usergrowth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMediaDataSync 媒体资源位投放效果数据回传
// taobao.usergrowth.ad.media.data.sync
//
// 创意维度广告效果数据回传
func TaobaoUsergrowthAdMediaDataSync(ctx context.Context, clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMediaDataSyncAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMediaDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
