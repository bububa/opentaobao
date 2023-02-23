package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundFillfeeAPIRequest 机票代理商】回填手续费 API请求
// taobao.alitrip.seller.refund.fillfee
//
// 回填手续费
type TaobaoAlitripSellerRefundFillfeeAPIRequest struct {
	model.Params
	// 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
	_feePriceMap string
	// 改签费用，格式：{detail_id:123,value:费用,金额单位分}
	_modifyFee string
	// 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
	_ticketPriceMap string
	// 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
	_upgradeFee string
	// 申请单ID
	_applyId int64
}

// NewTaobaoAlitripSellerRefundFillfeeRequest 初始化TaobaoAlitripSellerRefundFillfeeAPIRequest对象
func NewTaobaoAlitripSellerRefundFillfeeRequest() *TaobaoAlitripSellerRefundFillfeeAPIRequest {
	return &TaobaoAlitripSellerRefundFillfeeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.fillfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeePriceMap is FeePriceMap Setter
// 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeAPIRequest) SetFeePriceMap(_feePriceMap string) error {
	r._feePriceMap = _feePriceMap
	r.Set("fee_price_map", _feePriceMap)
	return nil
}

// GetFeePriceMap FeePriceMap Getter
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetFeePriceMap() string {
	return r._feePriceMap
}

// SetModifyFee is ModifyFee Setter
// 改签费用，格式：{detail_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeAPIRequest) SetModifyFee(_modifyFee string) error {
	r._modifyFee = _modifyFee
	r.Set("modify_fee", _modifyFee)
	return nil
}

// GetModifyFee ModifyFee Getter
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetModifyFee() string {
	return r._modifyFee
}

// SetTicketPriceMap is TicketPriceMap Setter
// 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeAPIRequest) SetTicketPriceMap(_ticketPriceMap string) error {
	r._ticketPriceMap = _ticketPriceMap
	r.Set("ticket_price_map", _ticketPriceMap)
	return nil
}

// GetTicketPriceMap TicketPriceMap Getter
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetTicketPriceMap() string {
	return r._ticketPriceMap
}

// SetUpgradeFee is UpgradeFee Setter
// 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeAPIRequest) SetUpgradeFee(_upgradeFee string) error {
	r._upgradeFee = _upgradeFee
	r.Set("upgrade_fee", _upgradeFee)
	return nil
}

// GetUpgradeFee UpgradeFee Getter
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetUpgradeFee() string {
	return r._upgradeFee
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundFillfeeAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerRefundFillfeeAPIRequest) GetApplyId() int64 {
	return r._applyId
}
