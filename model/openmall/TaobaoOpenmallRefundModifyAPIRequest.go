package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundmodifyAPIRequest 修改OpenMall退款申请 API请求
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
type TaobaoopenmallrefundmodifyAPIRequest struct {
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

// NewTaobaoopenmallrefundmodifyRequest 初始化TaobaoopenmallrefundmodifyAPIRequest对象
func NewTaobaoopenmallrefundmodifyRequest() *TaobaoopenmallrefundmodifyAPIRequest {
	return &TaobaoopenmallrefundmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundType is RefundType Setter
// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetRefundType(_refundType string) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRefundType() string {
	return r._refundType
}

// SetRefundDesc is RefundDesc Setter
// 买家的退货描述
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetRefundDesc(_refundDesc string) error {
	r._refundDesc = _refundDesc
	r.Set("refund_desc", _refundDesc)
	return nil
}

// GetRefundDesc RefundDesc Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRefundDesc() string {
	return r._refundDesc
}

// SetGoodsStatus is GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetGoodsStatus(_goodsStatus string) error {
	r._goodsStatus = _goodsStatus
	r.Set("goods_status", _goodsStatus)
	return nil
}

// GetGoodsStatus GoodsStatus Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetGoodsStatus() string {
	return r._goodsStatus
}

// SetRefundReason is RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetRefundReason(_refundReason string) error {
	r._refundReason = _refundReason
	r.Set("refund_reason", _refundReason)
	return nil
}

// GetRefundReason RefundReason Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRefundReason() string {
	return r._refundReason
}

// SetDistributor is Distributor Setter
// 分销者联盟身份
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundFee is RefundFee Setter
// 退款金额，分
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoopenmallrefundmodifyAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundmodifyAPIRequest) GetRefundId() int64 {
	return r._refundId
}
