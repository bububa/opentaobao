package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按编号查询采购申请/经销采购单 APIRequest
taobao.fenxiao.dealer.requisitionorder.query

按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
*/
type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    model.Params

    // 经销采购单编号。<br/>多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
    dealerOrderIds   []Number 

    // 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
    fields   string 

}

func NewTaobaoFenxiaoDealerRequisitionorderQueryRequest() *TaobaoFenxiaoDealerRequisitionorderQueryRequest{
    return &TaobaoFenxiaoDealerRequisitionorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.query"
}

func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetDealerOrderIds(dealerOrderIds []Number) error {
    r.dealerOrderIds = dealerOrderIds
    r.Set("dealer_order_ids", dealerOrderIds)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetDealerOrderIds() []Number {
    return r.dealerOrderIds
}

func (r *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetFields() string {
    return r.fields
}

