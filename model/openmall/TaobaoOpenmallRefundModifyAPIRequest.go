package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundModifyAPIRequest 修改OpenMall退款申请 API请求
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
type TaobaoOpenmallRefundModifyAPIRequest struct {
	model.Params
	// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
	_refundType string
	// 买家的退货描述
	_refundDesc string
	// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
	_goodsStatus string
	// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
	_refundReason string
	// 分销者联盟身份
	_distributor string
	// 退款金额，分
	_refundFee int64
	// 退款单ID
	_refundId int64
}

// NewTaobaoOpenmallRefundModifyRequest 初始化TaobaoOpenmallRefundModifyAPIRequest对象
func NewTaobaoOpenmallRefundModifyRequest() *TaobaoOpenmallRefundModifyAPIRequest {
	return &TaobaoOpenmallRefundModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundModifyAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundType is RefundType Setter
// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetRefundType(_refundType string) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRefundType() string {
	return r._refundType
}

// SetRefundDesc is RefundDesc Setter
// 买家的退货描述
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetRefundDesc(_refundDesc string) error {
	r._refundDesc = _refundDesc
	r.Set("refund_desc", _refundDesc)
	return nil
}

// GetRefundDesc RefundDesc Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRefundDesc() string {
	return r._refundDesc
}

// SetGoodsStatus is GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetGoodsStatus(_goodsStatus string) error {
	r._goodsStatus = _goodsStatus
	r.Set("goods_status", _goodsStatus)
	return nil
}

// GetGoodsStatus GoodsStatus Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetGoodsStatus() string {
	return r._goodsStatus
}

// SetRefundReason is RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetRefundReason(_refundReason string) error {
	r._refundReason = _refundReason
	r.Set("refund_reason", _refundReason)
	return nil
}

// GetRefundReason RefundReason Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRefundReason() string {
	return r._refundReason
}

// SetDistributor is Distributor Setter
// 分销者联盟身份
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundFee is RefundFee Setter
// 退款金额，分
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundModifyAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundModifyAPIRequest) GetRefundId() int64 {
	return r._refundId
}
