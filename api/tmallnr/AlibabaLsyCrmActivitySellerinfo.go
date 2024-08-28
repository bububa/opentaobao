package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivitySellerinfo 商家信息推送
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
func AlibabaLsyCrmActivitySellerinfo(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivitySellerinfoAPIRequest, resp *tmallnr.AlibabaLsyCrmActivitySellerinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
