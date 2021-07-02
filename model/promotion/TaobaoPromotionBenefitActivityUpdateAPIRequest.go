package promotion

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UpdateRequest Setter
// 修改关联的权益的活动请求
func (r *TaobaoPromotionBenefitActivityUpdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateBenefitActivityRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// Get UpdateRequest Getter
func (r TaobaoPromotionBenefitActivityUpdateAPIRequest) GetUpdateRequest() *UpdateBenefitActivityRequest {
	return r._updateRequest
}
