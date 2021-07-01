package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家收到的退款列表 API请求
taobao.refunds.receive.get

查询卖家收到的退款列表
*/
type TaobaoRefundsReceiveGetAPIRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
    _fields   []string
    // 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
    _status   string
    // 买家昵称
    _buyerNick   string
    // 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>
    _type   string
    // 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
    _startModified   string
    // 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
    _endModified   string
    // 页码。取值范围:大于零的整数; 默认值:1
    _pageNo   int64
    // 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    _pageSize   int64
    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
    _useHasNext   bool
}

// 初始化TaobaoRefundsReceiveGetAPIRequest对象
func NewTaobaoRefundsReceiveGetRequest() *TaobaoRefundsReceiveGetAPIRequest{
    return &TaobaoRefundsReceiveGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundsReceiveGetAPIRequest) GetApiMethodName() string {
    return "taobao.refunds.receive.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundsReceiveGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
func (r *TaobaoRefundsReceiveGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetFields() []string {
    return r._fields
}
// Status Setter
// 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
func (r *TaobaoRefundsReceiveGetAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetStatus() string {
    return r._status
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoRefundsReceiveGetAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Type Setter
// 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>
func (r *TaobaoRefundsReceiveGetAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetType() string {
    return r._type
}
// StartModified Setter
// 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
func (r *TaobaoRefundsReceiveGetAPIRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
func (r *TaobaoRefundsReceiveGetAPIRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetEndModified() string {
    return r._endModified
}
// PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaoRefundsReceiveGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
func (r *TaobaoRefundsReceiveGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
func (r *TaobaoRefundsReceiveGetAPIRequest) SetUseHasNext(_useHasNext bool) error {
    r._useHasNext = _useHasNext
    r.Set("use_has_next", _useHasNext)
    return nil
}

// UseHasNext Getter
func (r TaobaoRefundsReceiveGetAPIRequest) GetUseHasNext() bool {
    return r._useHasNext
}
