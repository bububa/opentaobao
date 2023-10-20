package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSpecialRefundsReceiveGetAPIRequest 特殊退款类型的纠纷单列表查询 API请求
// taobao.special.refunds.receive.get
//
// 特殊退款类型的纠纷单列表查询
type TaobaoSpecialRefundsReceiveGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
	_fields []string
	// 买家昵称
	_buyerNick string
	// 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
	_endModified string
	// 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
	_startModified string
	// 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
	_status string
	// 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>
	_type string
	// 买家openId
	_buyerOpenUid string
	// 页码。取值范围:大于零的整数; 默认值:1
	_pageNo int64
	// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
	_pageSize int64
	// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
	_useHasNext bool
}

// NewTaobaoSpecialRefundsReceiveGetRequest 初始化TaobaoSpecialRefundsReceiveGetAPIRequest对象
func NewTaobaoSpecialRefundsReceiveGetRequest() *TaobaoSpecialRefundsReceiveGetAPIRequest {
	return &TaobaoSpecialRefundsReceiveGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetApiMethodName() string {
	return "taobao.special.refunds.receive.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetEndModified is EndModified Setter
// 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetStartModified is StartModified Setter
// 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetStatus is Status Setter
// 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetStatus() string {
	return r._status
}

// SetType is Type Setter
// 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery这两种类型的数据，&lt;a href=&#34;http://open.taobao.com/doc/detail.htm?id=102855&#34; target=&#34;_blank&#34;&gt;查看可选值&lt;/a&gt;
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetType() string {
	return r._type
}

// SetBuyerOpenUid is BuyerOpenUid Setter
// 买家openId
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetBuyerOpenUid(_buyerOpenUid string) error {
	r._buyerOpenUid = _buyerOpenUid
	r.Set("buyer_open_uid", _buyerOpenUid)
	return nil
}

// GetBuyerOpenUid BuyerOpenUid Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetBuyerOpenUid() string {
	return r._buyerOpenUid
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetUseHasNext is UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。
func (r *TaobaoSpecialRefundsReceiveGetAPIRequest) SetUseHasNext(_useHasNext bool) error {
	r._useHasNext = _useHasNext
	r.Set("use_has_next", _useHasNext)
	return nil
}

// GetUseHasNext UseHasNext Getter
func (r TaobaoSpecialRefundsReceiveGetAPIRequest) GetUseHasNext() bool {
	return r._useHasNext
}
