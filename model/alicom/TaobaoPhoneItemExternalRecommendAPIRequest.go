package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneItemExternalRecommendAPIRequest 话费选品能力外放 API请求
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
type TaobaoPhoneItemExternalRecommendAPIRequest struct {
	model.Params
	// 选品请求
	_phoneDistributionRecommendReq *PhoneDistributionRecommendReq
}

// NewTaobaoPhoneItemExternalRecommendRequest 初始化TaobaoPhoneItemExternalRecommendAPIRequest对象
func NewTaobaoPhoneItemExternalRecommendRequest() *TaobaoPhoneItemExternalRecommendAPIRequest {
	return &TaobaoPhoneItemExternalRecommendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPhoneItemExternalRecommendAPIRequest) Reset() {
	r._phoneDistributionRecommendReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPhoneItemExternalRecommendAPIRequest) GetApiMethodName() string {
	return "taobao.phone.item.external.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPhoneItemExternalRecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPhoneItemExternalRecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneDistributionRecommendReq is PhoneDistributionRecommendReq Setter
// 选品请求
func (r *TaobaoPhoneItemExternalRecommendAPIRequest) SetPhoneDistributionRecommendReq(_phoneDistributionRecommendReq *PhoneDistributionRecommendReq) error {
	r._phoneDistributionRecommendReq = _phoneDistributionRecommendReq
	r.Set("phone_distribution_recommend_req", _phoneDistributionRecommendReq)
	return nil
}

// GetPhoneDistributionRecommendReq PhoneDistributionRecommendReq Getter
func (r TaobaoPhoneItemExternalRecommendAPIRequest) GetPhoneDistributionRecommendReq() *PhoneDistributionRecommendReq {
	return r._phoneDistributionRecommendReq
}

var poolTaobaoPhoneItemExternalRecommendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPhoneItemExternalRecommendRequest()
	},
}

// GetTaobaoPhoneItemExternalRecommendRequest 从 sync.Pool 获取 TaobaoPhoneItemExternalRecommendAPIRequest
func GetTaobaoPhoneItemExternalRecommendAPIRequest() *TaobaoPhoneItemExternalRecommendAPIRequest {
	return poolTaobaoPhoneItemExternalRecommendAPIRequest.Get().(*TaobaoPhoneItemExternalRecommendAPIRequest)
}

// ReleaseTaobaoPhoneItemExternalRecommendAPIRequest 将 TaobaoPhoneItemExternalRecommendAPIRequest 放入 sync.Pool
func ReleaseTaobaoPhoneItemExternalRecommendAPIRequest(v *TaobaoPhoneItemExternalRecommendAPIRequest) {
	v.Reset()
	poolTaobaoPhoneItemExternalRecommendAPIRequest.Put(v)
}
