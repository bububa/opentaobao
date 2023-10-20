package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivityupdateAPIRequest 修改关联的活动权益 API请求
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
type TaobaopromotionbenefitactivityupdateAPIRequest struct {
	model.Params
	// 修改关联的权益的活动请求
	_updateRequest *UpdateBenefitActivityRequest
}

// NewTaobaopromotionbenefitactivityupdateRequest 初始化TaobaopromotionbenefitactivityupdateAPIRequest对象
func NewTaobaopromotionbenefitactivityupdateRequest() *TaobaopromotionbenefitactivityupdateAPIRequest {
	return &TaobaopromotionbenefitactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 修改关联的权益的活动请求
func (r *TaobaopromotionbenefitactivityupdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateBenefitActivityRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r TaobaopromotionbenefitactivityupdateAPIRequest) GetUpdateRequest() *UpdateBenefitActivityRequest {
	return r._updateRequest
}
