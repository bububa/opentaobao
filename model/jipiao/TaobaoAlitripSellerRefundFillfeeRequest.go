package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票代理商】回填手续费 API请求
taobao.alitrip.seller.refund.fillfee

回填手续费
*/
type TaobaoAlitripSellerRefundFillfeeRequest struct {
    model.Params
    // 申请单ID
    _applyId   int64
    // 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
    _feePriceMap   string
    // 改签费用，格式：{detail_id:123,value:费用,金额单位分}
    _modifyFee   string
    // 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
    _ticketPriceMap   string
    // 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
    _upgradeFee   string
}

// 初始化TaobaoAlitripSellerRefundFillfeeRequest对象
func NewTaobaoAlitripSellerRefundFillfeeRequest() *TaobaoAlitripSellerRefundFillfeeRequest{
    return &TaobaoAlitripSellerRefundFillfeeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.fillfee"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApplyId() int64 {
    return r._applyId
}
// FeePriceMap Setter
// 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetFeePriceMap(_feePriceMap string) error {
    r._feePriceMap = _feePriceMap
    r.Set("fee_price_map", _feePriceMap)
    return nil
}

// FeePriceMap Getter
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetFeePriceMap() string {
    return r._feePriceMap
}
// ModifyFee Setter
// 改签费用，格式：{detail_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetModifyFee(_modifyFee string) error {
    r._modifyFee = _modifyFee
    r.Set("modify_fee", _modifyFee)
    return nil
}

// ModifyFee Getter
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetModifyFee() string {
    return r._modifyFee
}
// TicketPriceMap Setter
// 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetTicketPriceMap(_ticketPriceMap string) error {
    r._ticketPriceMap = _ticketPriceMap
    r.Set("ticket_price_map", _ticketPriceMap)
    return nil
}

// TicketPriceMap Getter
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetTicketPriceMap() string {
    return r._ticketPriceMap
}
// UpgradeFee Setter
// 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetUpgradeFee(_upgradeFee string) error {
    r._upgradeFee = _upgradeFee
    r.Set("upgrade_fee", _upgradeFee)
    return nil
}

// UpgradeFee Getter
func (r TaobaoAlitripSellerRefundFillfeeRequest) GetUpgradeFee() string {
    return r._upgradeFee
}
