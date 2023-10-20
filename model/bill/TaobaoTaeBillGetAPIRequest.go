package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaebillgetAPIRequest tae查询单笔账单明细 API请求
// taobao.tae.bill.get
//
// 查询单笔账单明细
type TaobaotaebillgetAPIRequest struct {
	model.Params
	// 传入需要返回的字段
	_fields []string
	// 账单编号
	_bid int64
	// 账单编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// NewTaobaotaebillgetRequest 初始化TaobaotaebillgetAPIRequest对象
func NewTaobaotaebillgetRequest() *TaobaotaebillgetAPIRequest {
	return &TaobaotaebillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaebillgetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaebillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaebillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 传入需要返回的字段
func (r *TaobaotaebillgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotaebillgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetBid is Bid Setter
// 账单编号
func (r *TaobaotaebillgetAPIRequest) SetBid(_bid int64) error {
	r._bid = _bid
	r.Set("bid", _bid)
	return nil
}

// GetBid Bid Getter
func (r TaobaotaebillgetAPIRequest) GetBid() int64 {
	return r._bid
}

// SetId is Id Setter
// 账单编号
func (r *TaobaotaebillgetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaotaebillgetAPIRequest) GetId() int64 {
	return r._id
}

// SetAccountId is AccountId Setter
// 虚拟账户科目编号
func (r *TaobaotaebillgetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaotaebillgetAPIRequest) GetAccountId() int64 {
	return r._accountId
}
