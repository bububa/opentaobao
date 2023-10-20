package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionBenefitActivitySendAPIRequest 手淘专用单用户发放接口 API请求
// taobao.mobile.promotion.benefit.activity.send
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
type TaobaoMobilePromotionBenefitActivitySendAPIRequest struct {
	model.Params
	// 单用户权益发放请求
	_singleBenefitRequest *SingleBenefitRequest
}

// NewTaobaoMobilePromotionBenefitActivitySendRequest 初始化TaobaoMobilePromotionBenefitActivitySendAPIRequest对象
func NewTaobaoMobilePromotionBenefitActivitySendRequest() *TaobaoMobilePromotionBenefitActivitySendAPIRequest {
	return &TaobaoMobilePromotionBenefitActivitySendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMobilePromotionBenefitActivitySendAPIRequest) Reset() {
	r._singleBenefitRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionBenefitActivitySendAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.benefit.activity.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionBenefitActivitySendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMobilePromotionBenefitActivitySendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSingleBenefitRequest is SingleBenefitRequest Setter
// 单用户权益发放请求
func (r *TaobaoMobilePromotionBenefitActivitySendAPIRequest) SetSingleBenefitRequest(_singleBenefitRequest *SingleBenefitRequest) error {
	r._singleBenefitRequest = _singleBenefitRequest
	r.Set("single_benefit_request", _singleBenefitRequest)
	return nil
}

// GetSingleBenefitRequest SingleBenefitRequest Getter
func (r TaobaoMobilePromotionBenefitActivitySendAPIRequest) GetSingleBenefitRequest() *SingleBenefitRequest {
	return r._singleBenefitRequest
}

var poolTaobaoMobilePromotionBenefitActivitySendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMobilePromotionBenefitActivitySendRequest()
	},
}

// GetTaobaoMobilePromotionBenefitActivitySendRequest 从 sync.Pool 获取 TaobaoMobilePromotionBenefitActivitySendAPIRequest
func GetTaobaoMobilePromotionBenefitActivitySendAPIRequest() *TaobaoMobilePromotionBenefitActivitySendAPIRequest {
	return poolTaobaoMobilePromotionBenefitActivitySendAPIRequest.Get().(*TaobaoMobilePromotionBenefitActivitySendAPIRequest)
}

// ReleaseTaobaoMobilePromotionBenefitActivitySendAPIRequest 将 TaobaoMobilePromotionBenefitActivitySendAPIRequest 放入 sync.Pool
func ReleaseTaobaoMobilePromotionBenefitActivitySendAPIRequest(v *TaobaoMobilePromotionBenefitActivitySendAPIRequest) {
	v.Reset()
	poolTaobaoMobilePromotionBenefitActivitySendAPIRequest.Put(v)
}
