package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobillaccountsgetAPIRequest 查询费用科目信息(限自研商家) API请求
// taobao.bill.accounts.get
//
// 查询费用账户信息
type TaobaobillaccountsgetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []string
}

// NewTaobaobillaccountsgetRequest 初始化TaobaobillaccountsgetAPIRequest对象
func NewTaobaobillaccountsgetRequest() *TaobaobillaccountsgetAPIRequest {
	return &TaobaobillaccountsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobillaccountsgetAPIRequest) GetApiMethodName() string {
	return "taobao.bill.accounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobillaccountsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobillaccountsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段
func (r *TaobaobillaccountsgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaobillaccountsgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetAids is Aids Setter
// 需要获取的科目ID
func (r *TaobaobillaccountsgetAPIRequest) SetAids(_aids []string) error {
	r._aids = _aids
	r.Set("aids", _aids)
	return nil
}

// GetAids Aids Getter
func (r TaobaobillaccountsgetAPIRequest) GetAids() []string {
	return r._aids
}
