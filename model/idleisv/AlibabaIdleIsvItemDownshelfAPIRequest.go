package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvItemDownshelfAPIRequest
服务商闲鱼商品下架 API请求
alibaba.idle.isv.item.downshelf

供外部服务商ISV进行闲鱼商品下架操作 */
type AlibabaIdleIsvItemDownshelfAPIRequest struct {
	model.Params
	// 返回结果result
	_param *IdleItemBaseApiDo
}

// NewAlibabaIdleIsvItemDownshelfRequest 初始化AlibabaIdleIsvItemDownshelfAPIRequest对象
func NewAlibabaIdleIsvItemDownshelfRequest() *AlibabaIdleIsvItemDownshelfAPIRequest {
	return &AlibabaIdleIsvItemDownshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 返回结果result
func (r *AlibabaIdleIsvItemDownshelfAPIRequest) SetParam(_param *IdleItemBaseApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetParam() *IdleItemBaseApiDo {
	return r._param
}
