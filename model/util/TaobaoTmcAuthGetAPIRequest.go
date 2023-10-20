package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcAuthGetAPIRequest) Reset() {
	r._group = ""
	r.Params.ToZero()
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

var poolTaobaoTmcAuthGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcAuthGetRequest()
	},
}

// GetTaobaoTmcAuthGetRequest 从 sync.Pool 获取 TaobaoTmcAuthGetAPIRequest
func GetTaobaoTmcAuthGetAPIRequest() *TaobaoTmcAuthGetAPIRequest {
	return poolTaobaoTmcAuthGetAPIRequest.Get().(*TaobaoTmcAuthGetAPIRequest)
}

// ReleaseTaobaoTmcAuthGetAPIRequest 将 TaobaoTmcAuthGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcAuthGetAPIRequest(v *TaobaoTmcAuthGetAPIRequest) {
	v.Reset()
	poolTaobaoTmcAuthGetAPIRequest.Put(v)
}
