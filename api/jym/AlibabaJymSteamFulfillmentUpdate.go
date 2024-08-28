package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamFulfillmentUpdate 交易猫Steam类目发履约态变更
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
func AlibabaJymSteamFulfillmentUpdate(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymSteamFulfillmentUpdateAPIRequest, resp *jym.AlibabaJymSteamFulfillmentUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
