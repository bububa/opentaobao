package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsordersgetAPIRequest 批量查询物流订单 API请求
// taobao.logistics.orders.get
//
// 批量查询物流订单。
type TaobaologisticsordersgetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br><br/>tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
	_fields string
	// 买家昵称
	_buyerNick string
	// 物流状态.查看数据结构 Shipping 中的status字段.
	_status string
	// 卖家是否发货.可选值:yes(是),no(否).如:yes
	_sellerConfirm string
	// 收货人姓名
	_receiverName string
	// 创建时间开始
	_startCreated string
	// 创建时间结束
	_endCreated string
	// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
	_freightPayer string
	// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
	_type string
	// 系统自动生成
	_ouid string
	// 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
	_tid int64
	// 页码.该字段没传 或 值<1 ,则默认page_no为1
	_pageNo int64
	// 每页条数.该字段没传 或 值<1 ,则默认page_size为40
	_pageSize int64
}

// NewTaobaologisticsordersgetRequest 初始化TaobaologisticsordersgetAPIRequest对象
func NewTaobaologisticsordersgetRequest() *TaobaologisticsordersgetAPIRequest {
	return &TaobaologisticsordersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsordersgetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.orders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsordersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsordersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: &lt;br&gt;&lt;br/&gt;tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；&lt;br&gt;多个字段之间用&#34;,&#34;分隔。如tid,seller_nick,buyer_nick,delivery_start。
func (r *TaobaologisticsordersgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaologisticsordersgetAPIRequest) GetFields() string {
	return r._fields
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaologisticsordersgetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaologisticsordersgetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetStatus is Status Setter
// 物流状态.查看数据结构 Shipping 中的status字段.
func (r *TaobaologisticsordersgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaologisticsordersgetAPIRequest) GetStatus() string {
	return r._status
}

// SetSellerConfirm is SellerConfirm Setter
// 卖家是否发货.可选值:yes(是),no(否).如:yes
func (r *TaobaologisticsordersgetAPIRequest) SetSellerConfirm(_sellerConfirm string) error {
	r._sellerConfirm = _sellerConfirm
	r.Set("seller_confirm", _sellerConfirm)
	return nil
}

// GetSellerConfirm SellerConfirm Getter
func (r TaobaologisticsordersgetAPIRequest) GetSellerConfirm() string {
	return r._sellerConfirm
}

// SetReceiverName is ReceiverName Setter
// 收货人姓名
func (r *TaobaologisticsordersgetAPIRequest) SetReceiverName(_receiverName string) error {
	r._receiverName = _receiverName
	r.Set("receiver_name", _receiverName)
	return nil
}

// GetReceiverName ReceiverName Getter
func (r TaobaologisticsordersgetAPIRequest) GetReceiverName() string {
	return r._receiverName
}

// SetStartCreated is StartCreated Setter
// 创建时间开始
func (r *TaobaologisticsordersgetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaologisticsordersgetAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetEndCreated is EndCreated Setter
// 创建时间结束
func (r *TaobaologisticsordersgetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaologisticsordersgetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetFreightPayer is FreightPayer Setter
// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
func (r *TaobaologisticsordersgetAPIRequest) SetFreightPayer(_freightPayer string) error {
	r._freightPayer = _freightPayer
	r.Set("freight_payer", _freightPayer)
	return nil
}

// GetFreightPayer FreightPayer Getter
func (r TaobaologisticsordersgetAPIRequest) GetFreightPayer() string {
	return r._freightPayer
}

// SetType is Type Setter
// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
func (r *TaobaologisticsordersgetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaologisticsordersgetAPIRequest) GetType() string {
	return r._type
}

// SetOuid is Ouid Setter
// 系统自动生成
func (r *TaobaologisticsordersgetAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaologisticsordersgetAPIRequest) GetOuid() string {
	return r._ouid
}

// SetTid is Tid Setter
// 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
func (r *TaobaologisticsordersgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticsordersgetAPIRequest) GetTid() int64 {
	return r._tid
}

// SetPageNo is PageNo Setter
// 页码.该字段没传 或 值&lt;1 ,则默认page_no为1
func (r *TaobaologisticsordersgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaologisticsordersgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数.该字段没传 或 值&lt;1 ,则默认page_size为40
func (r *TaobaologisticsordersgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaologisticsordersgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
