package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvItemEditAPIRequest
服务商闲鱼商品编辑 API请求
alibaba.idle.isv.item.edit

服务商ISV闲鱼商品编辑操作 */
type AlibabaIdleIsvItemEditAPIRequest struct {
	model.Params
	// 商品数据参数
	_param *IdleItemApiDo
}

// New
