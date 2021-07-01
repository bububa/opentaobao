package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBillAccountsGetAPIRequest
查询费用科目信息(限自研商家) API请求
taobao.bill.accounts.get

查询费用账户信息 */
type TaobaoBillAccountsGetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []int64
}

// NewTaobaoBillAccountsGetRequest 初始化TaobaoBillAccountsGetAPIRequest对象
func NewTaobaoBillAccountsGetRequest() *TaobaoBillAccountsGetAPIRequest {
	return &TaobaoBillAccountsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBillAccountsGetAPIRequest) GetApiMethodName() string {
	return "taobao.bill.accounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBillAccountsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Fields Setter
// 需要返回的字段
func (r *TaobaoBillAccountsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoBillAccountsGetAPIRequest) GetFields() []string {
	return r._fields
}

// Set is Aids Setter
// 需要获取的科目ID
func (r *TaobaoBillAccountsGetAPIRequest) SetAids(_aids []int64) error {
	r._aids = _aids
	r.Set("aids", _aids)
	return nil
}

// Get Aids Getter
func (r TaobaoBillAccountsGetAPIRequest) GetAids() []int64 {
	return r._aids
}
