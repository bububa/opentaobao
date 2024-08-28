package usergrowth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialDataSync 素材投放效果数据回传
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
func TaobaoUsergrowthAdMaterialDataSync(ctx context.Context, clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialDataSyncAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMaterialDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
