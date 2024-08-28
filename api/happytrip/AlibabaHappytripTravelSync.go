package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTravelSync 差旅申请单同步接口
// alibaba.happytrip.travel.sync
//
// 以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
func AlibabaHappytripTravelSync(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTravelSyncAPIRequest, resp *happytrip.AlibabaHappytripTravelSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
