package waybill

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillPrivacySubscriptionGetAPIRequest) Reset() {
	r.Params.ToZero()
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

var poolCainiaoWaybillPrivacySubscriptionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillPrivacySubscriptionGetRequest()
	},
}

// GetCainiaoWaybillPrivacySubscriptionGetRequest 从 sync.Pool 获取 CainiaoWaybillPrivacySubscriptionGetAPIRequest
func GetCainiaoWaybillPrivacySubscriptionGetAPIRequest() *CainiaoWaybillPrivacySubscriptionGetAPIRequest {
	return poolCainiaoWaybillPrivacySubscriptionGetAPIRequest.Get().(*CainiaoWaybillPrivacySubscriptionGetAPIRequest)
}

// ReleaseCainiaoWaybillPrivacySubscriptionGetAPIRequest 将 CainiaoWaybillPrivacySubscriptionGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillPrivacySubscriptionGetAPIRequest(v *CainiaoWaybillPrivacySubscriptionGetAPIRequest) {
	v.Reset()
	poolCainiaoWaybillPrivacySubscriptionGetAPIRequest.Put(v)
}
