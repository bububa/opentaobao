package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaebookbillgetAPIRequest tae查询单笔虚拟账户明细 API请求
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
type TaobaotaebookbillgetAPIRequest struct {
	model.Params
	// 需要返回的字段:参见BookBill结构体
	_fields []string
	// 虚拟账户流水编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// NewTaobaotaebookbillgetRequest 初始化TaobaotaebookbillgetAPIRequest对象
func NewTaobaotaebookbillgetRequest() *TaobaotaebookbillgetAPIRequest {
	return &TaobaotaebookbillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaebookbillgetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.book.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaebookbillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaebookbillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段:参见BookBill结构体
func (r *TaobaotaebookbillgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotaebookbillgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetId is Id Setter
// 虚拟账户流水编号
func (r *TaobaotaebookbillgetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaotaebookbillgetAPIRequest) GetId() int64 {
	return r._id
}

// SetAccountId is AccountId Setter
// 虚拟账户科目编号
func (r *TaobaotaebookbillgetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaotaebookbillgetAPIRequest) GetAccountId() int64 {
	return r._accountId
}
