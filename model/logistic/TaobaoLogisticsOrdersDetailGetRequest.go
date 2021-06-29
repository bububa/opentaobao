package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单,返回详细信息 API请求
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。
*/
type TaobaoLogisticsOrdersDetailGetRequest struct {
    model.Params
    // 需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔.
    fields   string
    // 交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息.
    tid   int64
    // 买家昵称
    buyerNick   string
    // 物流状态.可查看数据结构 Shipping 中的status字段.
    status   string
    // 卖家是否发货.可选值:yes(是),no(否).如:yes.
    sellerConfirm   string
    // 收货人姓名
    receiverName   string
    // 创建时间开始.格式:yyyy-MM-dd HH:mm:ss
    startCreated   string
    // 创建时间结束.格式:yyyy-MM-dd HH:mm:ss
    endCreated   string
    // 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
    freightPayer   string
    // 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
    type   string
    // 页码.该字段没传 或 值<1 ,则默认page_no为1
    pageNo   int64
    // 每页条数.该字段没传 或 值<1 ，则默认page_size为40
    pageSize   int64
}

// 初始化TaobaoLogisticsOrdersDetailGetRequest对象
func NewTaobaoLogisticsOrdersDetailGetRequest() *TaobaoLogisticsOrdersDetailGetRequest{
    return &TaobaoLogisticsOrdersDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrdersDetailGetRequest) GetApiMethodName() string {
    return "taobao.logistics.orders.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrdersDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔.
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetFields() string {
    return r.fields
}
// Tid Setter
// 交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息.
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetTid() int64 {
    return r.tid
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetBuyerNick() string {
    return r.buyerNick
}
// Status Setter
// 物流状态.可查看数据结构 Shipping 中的status字段.
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetStatus() string {
    return r.status
}
// SellerConfirm Setter
// 卖家是否发货.可选值:yes(是),no(否).如:yes.
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetSellerConfirm(sellerConfirm string) error {
    r.sellerConfirm = sellerConfirm
    r.Set("seller_confirm", sellerConfirm)
    return nil
}

// SellerConfirm Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetSellerConfirm() string {
    return r.sellerConfirm
}
// ReceiverName Setter
// 收货人姓名
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetReceiverName(receiverName string) error {
    r.receiverName = receiverName
    r.Set("receiver_name", receiverName)
    return nil
}

// ReceiverName Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetReceiverName() string {
    return r.receiverName
}
// StartCreated Setter
// 创建时间开始.格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetStartCreated() string {
    return r.startCreated
}
// EndCreated Setter
// 创建时间结束.格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetEndCreated() string {
    return r.endCreated
}
// FreightPayer Setter
// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetFreightPayer(freightPayer string) error {
    r.freightPayer = freightPayer
    r.Set("freight_payer", freightPayer)
    return nil
}

// FreightPayer Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetFreightPayer() string {
    return r.freightPayer
}
// Type Setter
// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetType() string {
    return r.type
}
// PageNo Setter
// 页码.该字段没传 或 值<1 ,则默认page_no为1
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数.该字段没传 或 值<1 ，则默认page_size为40
func (r *TaobaoLogisticsOrdersDetailGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoLogisticsOrdersDetailGetRequest) GetPageSize() int64 {
    return r.pageSize
}
