package bill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBillAccountsGetAPIRequest 查询费用科目信息(限自研商家) API请求
// taobao.bill.accounts.get
//
// 查询费用账户信息
type TaobaoBillAccountsGetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []string
}

// NewTaobaoBillAccountsGetRequest 初始化TaobaoBillAccountsGetAPIRequest对象
func NewTaobaoBillAccountsGetRequest() *TaobaoBillAccountsGetAPIRequest {
	return &TaobaoBillAccountsGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBillAccountsGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._aids = r._aids[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBillAccountsGetAPIRequest) GetApiMethodName() string {
	return "taobao.bill.accounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBillAccountsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBillAccountsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段
func (r *TaobaoBillAccountsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoBillAccountsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetAids is Aids Setter
// 需要获取的科目ID
func (r *TaobaoBillAccountsGetAPIRequest) SetAids(_aids []string) error {
	r._aids = _aids
	r.Set("aids", _aids)
	return nil
}

// GetAids Aids Getter
func (r TaobaoBillAccountsGetAPIRequest) GetAids() []string {
	return r._aids
}

var poolTaobaoBillAccountsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBillAccountsGetRequest()
	},
}

// GetTaobaoBillAccountsGetRequest 从 sync.Pool 获取 TaobaoBillAccountsGetAPIRequest
func GetTaobaoBillAccountsGetAPIRequest() *TaobaoBillAccountsGetAPIRequest {
	return poolTaobaoBillAccountsGetAPIRequest.Get().(*TaobaoBillAccountsGetAPIRequest)
}

// ReleaseTaobaoBillAccountsGetAPIRequest 将 TaobaoBillAccountsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBillAccountsGetAPIRequest(v *TaobaoBillAccountsGetAPIRequest) {
	v.Reset()
	poolTaobaoBillAccountsGetAPIRequest.Put(v)
}
