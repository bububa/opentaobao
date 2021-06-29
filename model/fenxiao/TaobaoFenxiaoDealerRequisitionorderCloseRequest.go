package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商关闭采购申请/经销采购单 API请求
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单
*/
type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    model.Params
    // 经销采购单编号
    _dealerOrderId   int64
    // 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
    _reason   int64
    // 关闭详细原因，字数5-200字
    _reasonDetail   string
}

// 初始化TaobaoFenxiaoDealerRequisitionorderCloseRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderCloseRequest() *TaobaoFenxiaoDealerRequisitionorderCloseRequest{
    return &TaobaoFenxiaoDealerRequisitionorderCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.close"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DealerOrderId Setter
// 经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetDealerOrderId(_dealerOrderId int64) error {
    r._dealerOrderId = _dealerOrderId
    r.Set("dealer_order_id", _dealerOrderId)
    return nil
}

// DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetDealerOrderId() int64 {
    return r._dealerOrderId
}
// Reason Setter
// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReason(_reason int64) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetReason() int64 {
    return r._reason
}
// ReasonDetail Setter
// 关闭详细原因，字数5-200字
func (r *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReasonDetail(_reasonDetail string) error {
    r._reasonDetail = _reasonDetail
    r.Set("reason_detail", _reasonDetail)
    return nil
}

// ReasonDetail Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseRequest) GetReasonDetail() string {
    return r._reasonDetail
}
