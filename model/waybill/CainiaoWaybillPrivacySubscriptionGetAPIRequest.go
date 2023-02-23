package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillPrivacySubscriptionGetAPIRequest 隐私面单商家订购查询 API请求
// cainiao.waybill.privacy.subscription.get
//
// ISV查询商家是否订购隐私面单
type CainiaoWaybillPrivacySubscriptionGetAPIRequest struct {
	model.Params
}

// NewCainiaoWaybillPrivacySubscriptionGetRequest 初始化CainiaoWaybillPrivacySubscriptionGetAPIRequest对象
func NewCainiaoWaybillPrivacySubscriptionGetRequest() *CainiaoWaybillPrivacySubscriptionGetAPIRequest {
	return &CainiaoWaybillPrivacySubscriptionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillPrivacySubscriptionGetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.privacy.subscription.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillPrivacySubscriptionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillPrivacySubscriptionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
