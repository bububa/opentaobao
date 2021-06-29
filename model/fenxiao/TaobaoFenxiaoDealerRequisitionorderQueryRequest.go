package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按编号查询采购申请/经销采购单 API请求
taobao.fenxiao.dealer.requisitionorder.query

按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
*/
type TaobaoFenxiaoDealerRequisitionorderQueryRequest struct {
    model.Params
    // 经销采购单编号。<br/>多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
    _dealerOrderIds   []int64
    // 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
    _fields   string
}

// 初始化TaobaoFenxiaoDealerRequisitionorderQueryRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderQueryRequest() *TaobaoFenxiaoDealerRequisitionorderQueryRequest{
    return &TaobaoFenxiaoDealerRequisitionorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DealerOrderIds Setter
// 经销采购单编号。<br/>多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。
func (r *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetDealerOrderIds(_dealerOrderIds []int64) error {
    r._dealerOrderIds = _dealerOrderIds
    r.Set("dealer_order_ids", _dealerOrderIds)
    return nil
}

// DealerOrderIds Getter
func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetDealerOrderIds() []int64 {
    return r._dealerOrderIds
}
// Fields Setter
// 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
func (r *TaobaoFenxiaoDealerRequisitionorderQueryRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoDealerRequisitionorderQueryRequest) GetFields() string {
    return r._fields
}
