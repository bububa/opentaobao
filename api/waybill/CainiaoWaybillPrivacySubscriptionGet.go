package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

/* CainiaoWaybillPrivacySubscriptionGet
隐私面单商家订购查询
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单 */
func CainiaoWaybillPrivacySubscriptionGet(clt *core.SDKClient, req *waybill.CainiaoWaybillPrivacySubscriptionGetAPIRequest, session string) (*waybill.CainiaoWaybillPrivacySubscriptionGetAPIResponse, error) {
	var resp waybill.CainiaoWaybillPrivacySubscriptionGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
