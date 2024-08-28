package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityPopupControl 营销弹屏疲劳度控制
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
func AlitripMerchantGalaxyActivityPopupControl(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityPopupControlAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityPopupControlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
