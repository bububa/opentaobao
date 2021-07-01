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

// New
