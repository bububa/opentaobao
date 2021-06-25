package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商关闭采购申请/经销采购单 APIRequest
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单
*/
type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    model.Params

    // 经销采购单编号
    dealerOrderId   int64 

    // 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
    reason   int64 

    // 关闭详细原因，字数5-200字
    reasonDetail   string 

}

func NewTaobaoFenxiaoDealerRequisitionorderCloseRequest() *TaobaoFenxiaoDealerRequisitionorderCloseRequest{
    return &TaobaoFenxiaoDealerRequisitionorderCloseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.close"
}

func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetDealerOrderId(dealerOrderId int64) error {
    r.dealerOrderId = dealerOrderId
    r.Set("dealer_order_id", dealerOrderId)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetDealerOrderId() int64 {
    return r.dealerOrderId
}

func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReason(reason int64) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetReason() int64 {
    return r.reason
}

func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReasonDetail(reasonDetail string) error {
    r.reasonDetail = reasonDetail
    r.Set("reason_detail", reasonDetail)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetReasonDetail() string {
    return r.reasonDetail
}

