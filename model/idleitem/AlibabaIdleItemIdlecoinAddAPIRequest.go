package idleitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleItemIdlecoinAddAPIRequest
免费送商品发送 API请求
alibaba.idle.item.idlecoin.add

免费送商品发布 */
type AlibabaIdleItemIdlecoinAddAPIRequest struct {
	model.Params
	// 免费送商品数据
	_idleCoinItemParam *IdleCoinItemApiDto
}

// New
