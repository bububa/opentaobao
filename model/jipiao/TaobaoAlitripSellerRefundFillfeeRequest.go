package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票代理商】回填手续费 APIRequest
taobao.alitrip.seller.refund.fillfee

回填手续费
*/
type TaobaoAlitripSellerRefundFillfeeRequest struct {
    model.Params

    // 申请单ID
    applyId   int64 

    // 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
    feePriceMap   string 

    // 改签费用，格式：{detail_id:123,value:费用,金额单位分}
    modifyFee   string 

    // 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
    ticketPriceMap   string 

    // 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
    upgradeFee   string 

}

func NewTaobaoAlitripSellerRefundFillfeeRequest() *TaobaoAlitripSellerRefundFillfeeRequest{
    return &TaobaoAlitripSellerRefundFillfeeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.fillfee"
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetFeePriceMap(feePriceMap string) error {
    r.feePriceMap = feePriceMap
    r.Set("fee_price_map", feePriceMap)
    return nil
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetFeePriceMap() string {
    return r.feePriceMap
}

func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetModifyFee(modifyFee string) error {
    r.modifyFee = modifyFee
    r.Set("modify_fee", modifyFee)
    return nil
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetModifyFee() string {
    return r.modifyFee
}

func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetTicketPriceMap(ticketPriceMap string) error {
    r.ticketPriceMap = ticketPriceMap
    r.Set("ticket_price_map", ticketPriceMap)
    return nil
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetTicketPriceMap() string {
    return r.ticketPriceMap
}

func (r *TaobaoAlitripSellerRefundFillfeeRequest) SetUpgradeFee(upgradeFee string) error {
    r.upgradeFee = upgradeFee
    r.Set("upgrade_fee", upgradeFee)
    return nil
}

func (r TaobaoAlitripSellerRefundFillfeeRequest) GetUpgradeFee() string {
    return r.upgradeFee
}

