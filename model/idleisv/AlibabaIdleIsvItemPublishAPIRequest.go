package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvItemPublishAPIRequest
服务商闲鱼商品发布 API请求
alibaba.idle.isv.item.publish

服务商ISV闲鱼商品发布 */
type AlibabaIdleIsvItemPublishAPIRequest struct {
	model.Params
	// 商品数据参数
	_itemParam *IdleItemApiDo
}

// New
