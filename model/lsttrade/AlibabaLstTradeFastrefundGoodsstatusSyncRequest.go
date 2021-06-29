package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家退款单商品状态同步 APIRequest
alibaba.lst.trade.fastrefund.goodsstatus.sync

卖家退款单商品状态同步
*/
type AlibabaLstTradeFastrefundGoodsstatusSyncRequest struct {
    model.Params

    // 主订单id
    mainOrderId   int64 

    // 退款单id
    refundId   string 

    // 未发货，枚举类型：UNSEND
    status   string 

}

func NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest() *AlibabaLstTradeFastrefundGoodsstatusSyncRequest{
    return &AlibabaLstTradeFastrefundGoodsstatusSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeFastrefundGoodsstatusSyncRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.fastrefund.goodsstatus.sync"
}

func (r AlibabaLstTradeFastrefundGoodsstatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeFastrefundGoodsstatusSyncRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r AlibabaLstTradeFastrefundGoodsstatusSyncRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *AlibabaLstTradeFastrefundGoodsstatusSyncRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaLstTradeFastrefundGoodsstatusSyncRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaLstTradeFastrefundGoodsstatusSyncRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaLstTradeFastrefundGoodsstatusSyncRequest) GetStatus() string {
    return r.status
}

