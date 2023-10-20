package bill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBookBillGetAPIRequest tae查询单笔虚拟账户明细 API请求
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetAPIRequest struct {
	model.Params
	// 需要返回的字段:参见BookBill结构体
	_fields []string
	// 虚拟账户流水编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// NewTaobaoTaeBookBillGetRequest 初始化TaobaoTaeBookBillGetAPIRequest对象
func NewTaobaoTaeBookBillGetRequest() *TaobaoTaeBookBillGetAPIRequest {
	return &TaobaoTaeBookBillGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaeBookBillGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._id = 0
	r._accountId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeBookBillGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.book.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeBookBillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaeBookBillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段:参见BookBill结构体
func (r *TaobaoTaeBookBillGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTaeBookBillGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetId is Id Setter
// 虚拟账户流水编号
func (r *TaobaoTaeBookBillGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoTaeBookBillGetAPIRequest) GetId() int64 {
	return r._id
}

// SetAccountId is AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBookBillGetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaoTaeBookBillGetAPIRequest) GetAccountId() int64 {
	return r._accountId
}

var poolTaobaoTaeBookBillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaeBookBillGetRequest()
	},
}

// GetTaobaoTaeBookBillGetRequest 从 sync.Pool 获取 TaobaoTaeBookBillGetAPIRequest
func GetTaobaoTaeBookBillGetAPIRequest() *TaobaoTaeBookBillGetAPIRequest {
	return poolTaobaoTaeBookBillGetAPIRequest.Get().(*TaobaoTaeBookBillGetAPIRequest)
}

// ReleaseTaobaoTaeBookBillGetAPIRequest 将 TaobaoTaeBookBillGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaeBookBillGetAPIRequest(v *TaobaoTaeBookBillGetAPIRequest) {
	v.Reset()
	poolTaobaoTaeBookBillGetAPIRequest.Put(v)
}
