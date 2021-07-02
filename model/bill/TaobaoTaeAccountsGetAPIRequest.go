package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeAccountsGetAPIRequest tae查询费用科目信息 API请求
// taobao.tae.accounts.get
//
// tae查询费用科目信息
type TaobaoTaeAccountsGetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []int64
}

// NewTaobaoTaeAccountsGetRequest 初始化TaobaoTaeAccountsGetAPIRequest对象
func NewTaobaoTaeAccountsGetRequest() *TaobaoTaeAccountsGetAPIRequest {
	return &TaobaoTaeAccountsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeAccountsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.accounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeAccountsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段
func (r *TaobaoTaeAccountsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTaeAccountsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetAids is Aids Setter
// 需要获取的科目ID
func (r *TaobaoTaeAccountsGetAPIRequest) SetAids(_aids []int64) error {
	r._aids = _aids
	r.Set("aids", _aids)
	return nil
}

// GetAids Aids Getter
func (r TaobaoTaeAccountsGetAPIRequest) GetAids() []int64 {
	return r._aids
}
