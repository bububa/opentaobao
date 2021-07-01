package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvItemQueryAPIRequest
服务商闲鱼商品查询 API请求
alibaba.idle.isv.item.query

服务商ISV闲鱼商品查询 */
type AlibabaIdleIsvItemQueryAPIRequest struct {
	model.Params
	// 商品查询参数
	_param *IdleItemBaseApiDo
}

// New
