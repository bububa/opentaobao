package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemdownshelfAPIRequest 服务商闲鱼商品下架 API请求
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
type AlibabaidleisvitemdownshelfAPIRequest struct {
	model.Params
	// 返回结果result
	_param *IdleItemBaseApiDo
}

// NewAlibabaidleisvitemdownshelfRequest 初始化AlibabaidleisvitemdownshelfAPIRequest对象
func NewAlibabaidleisvitemdownshelfRequest() *AlibabaidleisvitemdownshelfAPIRequest {
	return &AlibabaidleisvitemdownshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvitemdownshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvitemdownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvitemdownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 返回结果result
func (r *AlibabaidleisvitemdownshelfAPIRequest) SetParam(_param *IdleItemBaseApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaidleisvitemdownshelfAPIRequest) GetParam() *IdleItemBaseApiDo {
	return r._param
}
