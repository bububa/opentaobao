package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBillGetAPIRequest tae查询单笔账单明细 API请求
// taobao.tae.bill.get
//
// 查询单笔账单明细
type TaobaoTaeBillGetAPIRequest struct {
	model.Params
	// 账单编号
	_bid int64
	// 传入需要返回的字段
	_fields []string
	// 账单编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// NewTaobaoTaeBillGetRequest 初始化TaobaoTaeBillGetAPIRequest对象
func NewTaobaoTaeBillGetRequest() *TaobaoTaeBillGetAPIRequest {
	return &TaobaoTaeBillGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeBillGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeBillGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBid is Bid Setter
// 账单编号
func (r *TaobaoTaeBillGetAPIRequest) SetBid(_bid int64) error {
	r._bid = _bid
	r.Set("bid", _bid)
	return nil
}

// GetBid Bid Getter
func (r TaobaoTaeBillGetAPIRequest) GetBid() int64 {
	return r._bid
}

// SetFields is Fields Setter
// 传入需要返回的字段
func (r *TaobaoTaeBillGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTaeBillGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetId is Id Setter
// 账单编号
func (r *TaobaoTaeBillGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoTaeBillGetAPIRequest) GetId() int64 {
	return r._id
}

// SetAccountId is AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBillGetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaoTaeBillGetAPIRequest) GetAccountId() int64 {
	return r._accountId
}
