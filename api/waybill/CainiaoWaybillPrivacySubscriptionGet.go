package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillPrivacySubscriptionGet 隐私面单商家订购查询
// cainiao.waybill.privacy.subscription.get
//
// ISV查询商家是否订购隐私面单
func CainiaoWaybillPrivacySubscriptionGet(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillPrivacySubscriptionGetAPIRequest, resp *waybill.CainiaoWaybillPrivacySubscriptionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
