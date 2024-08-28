package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdatestatus 更新渠道店营业状态
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
func AlibabaAelophyShopUpdatestatus(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdatestatusAPIRequest, resp *wdk.AlibabaAelophyShopUpdatestatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
