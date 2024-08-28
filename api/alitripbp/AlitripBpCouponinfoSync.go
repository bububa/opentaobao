package alitripbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

// AlitripBpCouponinfoSync 飞猪广告券信息同步接口
// alitrip.bp.couponinfo.sync
//
// 飞猪商业化券信息同步
func AlitripBpCouponinfoSync(ctx context.Context, clt *core.SDKClient, req *alitripbp.AlitripBpCouponinfoSyncAPIRequest, resp *alitripbp.AlitripBpCouponinfoSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
