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
func (r TaobaoTmcAuthGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Group Setter
// tmc组名
func (r *TaobaoTmcAuthGetAPIRequest) SetGroup(_group string) error {
	r._group = _group
	r.Set("group", _group)
	return nil
}

// Get Group Getter
func (r TaobaoTmcAuthGetAPIRequest) GetGroup() string {
	return r._group
}
