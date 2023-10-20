package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundcreateAPIRequest 创建OpenMall退款单 API请求
// taobao.openmall.refund.create
//
// 创建OpenMall退款单
// 如存在未完结的退款单，则返回该退款单ID
type TaobaoopenmallrefundcreateAPIRequest struct {
	model.Params
	// 分销者联盟身份
	_distributor string
	// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
	_goodsStatus string
	// 买家的退货描述
	_refundDesc string
	// 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
	_refundReason string
	// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
	_refundType string
	// 退款金额，分
	_refundFee int64
	// 订单号
	_tid int64
}

// NewTaobaoopenmallrefundcreateRequest 初始化TaobaoopenmallrefundcreateAPIRequest对象
func NewTaobaoopenmallrefundcreateRequest() *TaobaoopenmallrefundcreateAPIRequest {
	return &TaobaoopenmallrefundcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundcreateAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者联盟身份
func (r *TaobaoopenmallrefundcreateAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetGoodsStatus is GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoopenmallrefundcreateAPIRequest) SetGoodsStatus(_goodsStatus string) error {
	r._goodsStatus = _goodsStatus
	r.Set("goods_status", _goodsStatus)
	return nil
}

// GetGoodsStatus GoodsStatus Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetGoodsStatus() string {
	return r._goodsStatus
}

// SetRefundDesc is RefundDesc Setter
// 买家的退货描述
func (r *TaobaoopenmallrefundcreateAPIRequest) SetRefundDesc(_refundDesc string) error {
	r._refundDesc = _refundDesc
	r.Set("refund_desc", _refundDesc)
	return nil
}

// GetRefundDesc RefundDesc Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetRefundDesc() string {
	return r._refundDesc
}

// SetRefundReason is RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoopenmallrefundcreateAPIRequest) SetRefundReason(_refundReason string) error {
	r._refundReason = _refundReason
	r.Set("refund_reason", _refundReason)
	return nil
}

// GetRefundReason RefundReason Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetRefundReason() string {
	return r._refundReason
}

// SetRefundType is RefundType Setter
// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
func (r *TaobaoopenmallrefundcreateAPIRequest) SetRefundType(_refundType string) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetRefundType() string {
	return r._refundType
}

// SetRefundFee is RefundFee Setter
// 退款金额，分
func (r *TaobaoopenmallrefundcreateAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetTid is Tid Setter
// 订单号
func (r *TaobaoopenmallrefundcreateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenmallrefundcreateAPIRequest) GetTid() int64 {
	return r._tid
}
