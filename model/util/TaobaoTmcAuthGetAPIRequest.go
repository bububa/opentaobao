package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcAuthGetAPIRequest TMC授权token API请求
// taobao.tmc.auth.get
//
// TMC连接授权Token
type TaobaoTmcAuthGetAPIRequest struct {
	model.Params
	// tmc组名
	_group string
}

// NewTaobaoTmcAuthGetRequest 初始化TaobaoTmcAuthGetAPIRequest对象
func NewTaobaoTmcAuthGetRequest() *TaobaoTmcAuthGetAPIRequest {
	return &TaobaoTmcAuthGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcAuthGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.auth.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcAuthGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcAuthGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroup is Group Setter
// tmc组名
func (r *TaobaoTmcAuthGetAPIRequest) SetGroup(_group string) error {
	r._group = _group
	r.Set("group", _group)
	return nil
}

// GetGroup Group Getter
func (r TaobaoTmcAuthGetAPIRequest) GetGroup() string {
	return r._group
}
