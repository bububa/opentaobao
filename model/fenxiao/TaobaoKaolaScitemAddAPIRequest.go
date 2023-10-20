package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokaolascitemaddAPIRequest 考拉货品新增接口 API请求
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
type TaobaokaolascitemaddAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 新增选项
	_option *AddCnskuOption
}

// NewTaobaokaolascitemaddRequest 初始化TaobaokaolascitemaddAPIRequest对象
func NewTaobaokaolascitemaddRequest() *TaobaokaolascitemaddAPIRequest {
	return &TaobaokaolascitemaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokaolascitemaddAPIRequest) GetApiMethodName() string {
	return "taobao.kaola.scitem.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokaolascitemaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokaolascitemaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnsku is Cnsku Setter
// 待新增的货品
func (r *TaobaokaolascitemaddAPIRequest) SetCnsku(_cnsku *CnskuDto) error {
	r._cnsku = _cnsku
	r.Set("cnsku", _cnsku)
	return nil
}

// GetCnsku Cnsku Getter
func (r TaobaokaolascitemaddAPIRequest) GetCnsku() *CnskuDto {
	return r._cnsku
}

// SetOption is Option Setter
// 新增选项
func (r *TaobaokaolascitemaddAPIRequest) SetOption(_option *AddCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r TaobaokaolascitemaddAPIRequest) GetOption() *AddCnskuOption {
	return r._option
}
