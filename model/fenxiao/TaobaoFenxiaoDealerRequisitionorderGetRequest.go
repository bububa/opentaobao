package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购申请/经销采购单 API请求
taobao.fenxiao.dealer.requisitionorder.get

批量查询采购申请/经销采购单，目前支持供应商和分销商查询
*/
type TaobaoFenxiaoDealerRequisitionorderGetRequest struct {
    model.Params
    // 采购申请/经销采购单最早修改时间
    startDate   string
    // 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天
    endDate   string
    // 页码（大于0的整数。无值或小于1的值按默认值1计）
    pageNo   int64
    // 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
    pageSize   int64
    // 采购申请/经销采购单状态。<br/>0：全部状态。<br/>1：分销商提交申请，待供应商审核。<br/>2：供应商驳回申请，待分销商确认。<br/>3：供应商修改后，待分销商确认。<br/>4：分销商拒绝修改，待供应商再审核。<br/>5：审核通过下单成功，待分销商付款。<br/>7：付款成功，待供应商发货。<br/>8：供应商发货，待分销商收货。<br/>9：分销商收货，交易成功。<br/>10：采购申请/经销采购单关闭。<br/><br/>注：无值按默认值0计，超出状态范围返回错误信息。
    orderStatus   int64
    // 多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
    fields   string
    // 查询者自己在所要查询的采购申请/经销采购单中的身份。<br/>1：供应商。<br/>2：分销商。<br/>注：填写其他值当做错误处理。
    identity   int64
}

// 初始化TaobaoFenxiaoDealerRequisitionorderGetRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderGetRequest() *TaobaoFenxiaoDealerRequisitionorderGetRequest{
    return &TaobaoFenxiaoDealerRequisitionorderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 采购申请/经销采购单最早修改时间
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetEndDate() string {
    return r.endDate
}
// PageNo Setter
// 页码（大于0的整数。无值或小于1的值按默认值1计）
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// OrderStatus Setter
// 采购申请/经销采购单状态。<br/>0：全部状态。<br/>1：分销商提交申请，待供应商审核。<br/>2：供应商驳回申请，待分销商确认。<br/>3：供应商修改后，待分销商确认。<br/>4：分销商拒绝修改，待供应商再审核。<br/>5：审核通过下单成功，待分销商付款。<br/>7：付款成功，待供应商发货。<br/>8：供应商发货，待分销商收货。<br/>9：分销商收货，交易成功。<br/>10：采购申请/经销采购单关闭。<br/><br/>注：无值按默认值0计，超出状态范围返回错误信息。
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetOrderStatus() int64 {
    return r.orderStatus
}
// Fields Setter
// 多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetFields() string {
    return r.fields
}
// Identity Setter
// 查询者自己在所要查询的采购申请/经销采购单中的身份。<br/>1：供应商。<br/>2：分销商。<br/>注：填写其他值当做错误处理。
func (r *TaobaoFenxiaoDealerRequisitionorderGetRequest) SetIdentity(identity int64) error {
    r.identity = identity
    r.Set("identity", identity)
    return nil
}

// Identity Getter
func (r TaobaoFenxiaoDealerRequisitionorderGetRequest) GetIdentity() int64 {
    return r.identity
}
