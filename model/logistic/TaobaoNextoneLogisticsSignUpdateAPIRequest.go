package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNextoneLogisticsSignUpdateAPIRequest AG物流签收状态写接口 API请求
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
type TaobaoNextoneLogisticsSignUpdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 货物签收状态
	_signStatus int64
}

// NewTaobaoNextoneLogisticsSignUpdateRequest 初始化TaobaoNextoneLogisticsSignUpdateAPIRequest对象
func NewTaobaoNextoneLogisticsSignUpdateRequest() *TaobaoNextoneLogisticsSignUpdateAPIRequest {
	return &TaobaoNextoneLogisticsSignUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.nextone.logistics.sign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefundId Setter
// 退款编号
func (r *TaobaoNextoneLogisticsSignUpdateAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// Set is SignStatus Setter
// 货物签收状态
func (r *TaobaoNextoneLogisticsSignUpdateAPIRequest) SetSignStatus(_signStatus int64) error {
	r._signStatus = _signStatus
	r.Set("sign_status", _signStatus)
	return nil
}

// Get SignStatus Getter
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetSignStatus() int64 {
	return r._signStatus
}
