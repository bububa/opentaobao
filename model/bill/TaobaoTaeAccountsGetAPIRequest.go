package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaeaccountsgetAPIRequest tae查询费用科目信息 API请求
// taobao.tae.accounts.get
//
// tae查询费用科目信息
type TaobaotaeaccountsgetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []string
}

// NewTaobaotaeaccountsgetRequest 初始化TaobaotaeaccountsgetAPIRequest对象
func NewTaobaotaeaccountsgetRequest() *TaobaotaeaccountsgetAPIRequest {
	return &TaobaotaeaccountsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaeaccountsgetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.accounts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaeaccountsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaeaccountsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段
func (r *TaobaotaeaccountsgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotaeaccountsgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetAids is Aids Setter
// 需要获取的科目ID
func (r *TaobaotaeaccountsgetAPIRequest) SetAids(_aids []string) error {
	r._aids = _aids
	r.Set("aids", _aids)
	return nil
}

// GetAids Aids Getter
func (r TaobaotaeaccountsgetAPIRequest) GetAids() []string {
	return r._aids
}
