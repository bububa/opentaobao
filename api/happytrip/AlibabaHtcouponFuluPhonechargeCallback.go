package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHtcouponFuluPhonechargeCallback 话费充值回调
// alibaba.htcoupon.fulu.phonecharge.callback
//
// 话费充值为异步操作，此接口用于充值成功后，供应商回调。
func AlibabaHtcouponFuluPhonechargeCallback(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHtcouponFuluPhonechargeCallbackAPIRequest, resp *happytrip.AlibabaHtcouponFuluPhonechargeCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
