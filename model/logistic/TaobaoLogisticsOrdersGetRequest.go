package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单 API请求
taobao.logistics.orders.get

批量查询物流订单。
*/
type TaobaoLogisticsOrdersGetRequest struct {
    model.Params
    // 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br><br/>tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
    _fields   string
    // 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
    _tid   int64
    // 买家昵称
    _buyerNick   string
    // 物流状态.查看数据结构 Shipping 中的status字段.
    _status   string
    // 卖家是否发货.可选值:yes(是),no(否).如:yes
    _sellerConfirm   string
    // 收货人姓名
    _receiverName   string
    // 创建时间开始
    _startCreated   string
    // 创建时间结束
    _endCreated   string
    // 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
    _freightPayer   string
    // 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
    _type   string
    // 页码.该字段没传 或 值<1 ,则默认page_no为1
    _pageNo   int64
    // 每页条数.该字段没传 或 值<1 ,则默认page_size为40
    _pageSize   int64
}

// 初始化TaobaoLogisticsOrdersGetRequest对象
func NewTaobaoLogisticsOrdersGetRequest() *TaobaoLogisticsOrdersGetRequest{
    return &TaobaoLogisticsOrdersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrdersGetRequest) GetApiMethodName() string {
    return "taobao.logistics.orders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br><br/>tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
func (r *TaobaoLogisticsOrdersGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoLogisticsOrdersGetRequest) GetFields() string {
    return r._fields
}
// Tid Setter
// 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
func (r *TaobaoLogisticsOrdersGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOrdersGetRequest) GetTid() int64 {
    return r._tid
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoLogisticsOrdersGetRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoLogisticsOrdersGetRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Status Setter
// 物流状态.查看数据结构 Shipping 中的status字段.
func (r *TaobaoLogisticsOrdersGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoLogisticsOrdersGetRequest) GetStatus() string {
    return r._status
}
// SellerConfirm Setter
// 卖家是否发货.可选值:yes(是),no(否).如:yes
func (r *TaobaoLogisticsOrdersGetRequest) SetSellerConfirm(_sellerConfirm string) error {
    r._sellerConfirm = _sellerConfirm
    r.Set("seller_confirm", _sellerConfirm)
    return nil
}

// SellerConfirm Getter
func (r TaobaoLogisticsOrdersGetRequest) GetSellerConfirm() string {
    return r._sellerConfirm
}
// ReceiverName Setter
// 收货人姓名
func (r *TaobaoLogisticsOrdersGetRequest) SetReceiverName(_receiverName string) error {
    r._receiverName = _receiverName
    r.Set("receiver_name", _receiverName)
    return nil
}

// ReceiverName Getter
func (r TaobaoLogisticsOrdersGetRequest) GetReceiverName() string {
    return r._receiverName
}
// StartCreated Setter
// 创建时间开始
func (r *TaobaoLogisticsOrdersGetRequest) SetStartCreated(_startCreated string) error {
    r._startCreated = _startCreated
    r.Set("start_created", _startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoLogisticsOrdersGetRequest) GetStartCreated() string {
    return r._startCreated
}
// EndCreated Setter
// 创建时间结束
func (r *TaobaoLogisticsOrdersGetRequest) SetEndCreated(_endCreated string) error {
    r._endCreated = _endCreated
    r.Set("end_created", _endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoLogisticsOrdersGetRequest) GetEndCreated() string {
    return r._endCreated
}
// FreightPayer Setter
// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
func (r *TaobaoLogisticsOrdersGetRequest) SetFreightPayer(_freightPayer string) error {
    r._freightPayer = _freightPayer
    r.Set("freight_payer", _freightPayer)
    return nil
}

// FreightPayer Getter
func (r TaobaoLogisticsOrdersGetRequest) GetFreightPayer() string {
    return r._freightPayer
}
// Type Setter
// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
func (r *TaobaoLogisticsOrdersGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoLogisticsOrdersGetRequest) GetType() string {
    return r._type
}
// PageNo Setter
// 页码.该字段没传 或 值<1 ,则默认page_no为1
func (r *TaobaoLogisticsOrdersGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoLogisticsOrdersGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数.该字段没传 或 值<1 ,则默认page_size为40
func (r *TaobaoLogisticsOrdersGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoLogisticsOrdersGetRequest) GetPageSize() int64 {
    return r._pageSize
}
