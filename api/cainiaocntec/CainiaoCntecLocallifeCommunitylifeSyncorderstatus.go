package cainiaocntec

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeSyncorderstatus 订单状态推送
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
func CainiaoCntecLocallifeCommunitylifeSyncorderstatus(ctx context.Context, clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest, resp *cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
