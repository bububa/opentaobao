package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemQueryAPIRequest 服务商闲鱼商品查询 API请求
// alibaba.idle.isv.item.query
//
// 服务商ISV闲鱼商品查询
type AlibabaIdleIsvItemQueryAPIRequest struct {
	model.Params
	// 商品查询参数
	_param *IdleItemBaseApiDo
}

// NewAlibabaIdleIsvItemQueryRequest 初始化AlibabaIdleIsvItemQueryAPIRequest对象
func NewAlibabaIdleIsvItemQueryRequest() *AlibabaIdleIsvItemQueryAPIRequest {
	return &AlibabaIdleIsvItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 商品查询参数
func (r *AlibabaIdleIsvItemQueryAPIRequest) SetParam(_param *IdleItemBaseApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleIsvItemQueryAPIRequest) GetParam() *IdleItemBaseApiDo {
	return r._param
}
