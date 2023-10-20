package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemeditAPIRequest 服务商闲鱼商品编辑 API请求
// alibaba.idle.isv.item.edit
//
// 服务商ISV闲鱼商品编辑操作
type AlibabaidleisvitemeditAPIRequest struct {
	model.Params
	// 商品数据参数
	_param *IdleItemApiDo
}

// NewAlibabaidleisvitemeditRequest 初始化AlibabaidleisvitemeditAPIRequest对象
func NewAlibabaidleisvitemeditRequest() *AlibabaidleisvitemeditAPIRequest {
	return &AlibabaidleisvitemeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvitemeditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvitemeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvitemeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 商品数据参数
func (r *AlibabaidleisvitemeditAPIRequest) SetParam(_param *IdleItemApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaidleisvitemeditAPIRequest) GetParam() *IdleItemApiDo {
	return r._param
}
