package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商通过采购申请/经销采购单申请 API请求
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核
*/
type TaobaoFenxiaoDealerRequisitionorderAgreeRequest struct {
    model.Params
    // 采购申请/经销采购单编号
    dealerOrderId   int64
}

// 初始化TaobaoFenxiaoDealerRequisitionorderAgreeRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest() *TaobaoFenxiaoDealerRequisitionorderAgreeRequest{
    return &TaobaoFenxiaoDealerRequisitionorderAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.agree"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DealerOrderId Setter
// 采购申请/经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) SetDealerOrderId(dealerOrderId int64) error {
    r.dealerOrderId = dealerOrderId
    r.Set("dealer_order_id", dealerOrderId)
    return nil
}

// DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderAgreeRequest) GetDealerOrderId() int64 {
    return r.dealerOrderId
}
