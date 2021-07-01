package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcAuthGetAPIRequest
TMC授权token API请求
taobao.tmc.auth.get

TMC连接授权Token */
type TaobaoTmcAuthGetAPIRequest struct {
	model.Params
	// tmc组名
	_group string
}

// New
