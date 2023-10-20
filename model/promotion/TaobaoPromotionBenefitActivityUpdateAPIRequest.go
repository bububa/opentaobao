package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityUpdateAPIRequest 修改关联的活动权益 API请求
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
type TaobaoPromotionBenefitActivityUpdateAPIRequest struct {
	model.Params
	// 修改关联的权益的活动请求
	_updateRequest *UpdateBenefitActivityRequest
}

// NewTaobaoPromotionBenefitActivityUpdateRequest 初始化TaobaoPromotionBenefitActivityUpdateAPIRequest对象
func NewTaobaoPromotionBenefitActivityUpdateRequest() *TaobaoPromotionBenefitActivityUpdateAPIRequest {
	return &TaobaoPromotionBenefitActivityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionBenefitActivityUpdateAPIRequest) Reset() {
	r._updateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 修改关联的权益的活动请求
func (r *TaobaoPromotionBenefitActivityUpdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateBenefitActivityRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetUpdateRequest() *UpdateBenefitActivityRequest {
	return r._updateRequest
}

var poolTaobaoPromotionBenefitActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionBenefitActivityUpdateRequest()
	},
}

// GetTaobaoPromotionBenefitActivityUpdateRequest 从 sync.Pool 获取 TaobaoPromotionBenefitActivityUpdateAPIRequest
func GetTaobaoPromotionBenefitActivityUpdateAPIRequest() *TaobaoPromotionBenefitActivityUpdateAPIRequest {
	return poolTaobaoPromotionBenefitActivityUpdateAPIRequest.Get().(*TaobaoPromotionBenefitActivityUpdateAPIRequest)
}

// ReleaseTaobaoPromotionBenefitActivityUpdateAPIRequest 将 TaobaoPromotionBenefitActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityUpdateAPIRequest(v *TaobaoPromotionBenefitActivityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityUpdateAPIRequest.Put(v)
}
