package bill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBillGetAPIRequest tae查询单笔账单明细 API请求
// taobao.tae.bill.get
//
// 查询单笔账单明细
type TaobaoTaeBillGetAPIRequest struct {
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

// NewTaobaoTaeBillGetRequest 初始化TaobaoTaeBillGetAPIRequest对象
func NewTaobaoTaeBillGetRequest() *TaobaoTaeBillGetAPIRequest {
	return &TaobaoTaeBillGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaeBillGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._bid = 0
	r._id = 0
	r._accountId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeBillGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeBillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaeBillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTaeBillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaeBillGetRequest()
	},
}

// GetTaobaoTaeBillGetRequest 从 sync.Pool 获取 TaobaoTaeBillGetAPIRequest
func GetTaobaoTaeBillGetAPIRequest() *TaobaoTaeBillGetAPIRequest {
	return poolTaobaoTaeBillGetAPIRequest.Get().(*TaobaoTaeBillGetAPIRequest)
}

// ReleaseTaobaoTaeBillGetAPIRequest 将 TaobaoTaeBillGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaeBillGetAPIRequest(v *TaobaoTaeBillGetAPIRequest) {
	v.Reset()
	poolTaobaoTaeBillGetAPIRequest.Put(v)
}
