package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家回填物流信息 API请求
taobao.rp.returngoods.refill

卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
*/
type TaobaoRpReturngoodsRefillRequest struct {
    model.Params
    // 退款单编号
    refundId   int64
    // 退款阶段，可选值：售中：onsale，售后：aftersale
    refundPhase   string
    // 物流公司运单号
    logisticsWaybillNo   string
    // 物流公司编号
    logisticsCompanyCode   string
}

// 初始化TaobaoRpReturngoodsRefillRequest对象
func NewTaobaoRpReturngoodsRefillRequest() *TaobaoRpReturngoodsRefillRequest{
    return &TaobaoRpReturngoodsRefillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefillRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.refill"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单编号
func (r *TaobaoRpReturngoodsRefillRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsRefillRequest) GetRefundId() int64 {
    return r.refundId
}
// RefundPhase Setter
// 退款阶段，可选值：售中：onsale，售后：aftersale
func (r *TaobaoRpReturngoodsRefillRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsRefillRequest) GetRefundPhase() string {
    return r.refundPhase
}
// LogisticsWaybillNo Setter
// 物流公司运单号
func (r *TaobaoRpReturngoodsRefillRequest) SetLogisticsWaybillNo(logisticsWaybillNo string) error {
    r.logisticsWaybillNo = logisticsWaybillNo
    r.Set("logistics_waybill_no", logisticsWaybillNo)
    return nil
}

// LogisticsWaybillNo Getter
func (r TaobaoRpReturngoodsRefillRequest) GetLogisticsWaybillNo() string {
    return r.logisticsWaybillNo
}
// LogisticsCompanyCode Setter
// 物流公司编号
func (r *TaobaoRpReturngoodsRefillRequest) SetLogisticsCompanyCode(logisticsCompanyCode string) error {
    r.logisticsCompanyCode = logisticsCompanyCode
    r.Set("logistics_company_code", logisticsCompanyCode)
    return nil
}

// LogisticsCompanyCode Getter
func (r TaobaoRpReturngoodsRefillRequest) GetLogisticsCompanyCode() string {
    return r.logisticsCompanyCode
}
