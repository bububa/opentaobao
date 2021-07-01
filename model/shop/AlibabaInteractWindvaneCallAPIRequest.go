package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractWindvaneCallAPIRequest
Weex页面调用windvane API请求
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据 */
type AlibabaInteractWindvaneCallAPIRequest struct {
	model.Params
	// 客户端鉴权使用，实际不会发送或接收数据
	_unNamed string
}

// New
