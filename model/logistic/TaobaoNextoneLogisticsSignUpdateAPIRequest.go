package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaonextonelogisticssignupdateAPIRequest AG物流签收状态写接口 API请求
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
type TaobaonextonelogisticssignupdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 货物签收状态
	_signStatus int64
}

// NewTaobaonextonelogisticssignupdateRequest 初始化TaobaonextonelogisticssignupdateAPIRequest对象
func NewTaobaonextonelogisticssignupdateRequest() *TaobaonextonelogisticssignupdateAPIRequest {
	return &TaobaonextonelogisticssignupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaonextonelogisticssignupdateAPIRequest) GetApiMethodName() string {
	return "taobao.nextone.logistics.sign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaonextonelogisticssignupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaonextonelogisticssignupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaonextonelogisticssignupdateAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaonextonelogisticssignupdateAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetSignStatus is SignStatus Setter
// 货物签收状态
func (r *TaobaonextonelogisticssignupdateAPIRequest) SetSignStatus(_signStatus int64) error {
	r._signStatus = _signStatus
	r.Set("sign_status", _signStatus)
	return nil
}

// GetSignStatus SignStatus Getter
func (r TaobaonextonelogisticssignupdateAPIRequest) GetSignStatus() int64 {
	return r._signStatus
}
