package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveTaskRegisterAPIRequest
互动任务开通接口 API请求
taobao.jst.interactive.task.register

调用互动任务开通接口为小程序开通互动任务 */
type TaobaoJstInteractiveTaskRegisterAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// New
