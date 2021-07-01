package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillPrivacySubscriptionGetAPIRequest
隐私面单商家订购查询 API请求
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单 */
type CainiaoWaybillPrivacySubscriptionGetAPIRequest struct {
	model.Params
}

// New
