package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenlinkSessionGetAPIRequest
获取授权session信息 API请求
taobao.openlink.session.get

帮助第三方isv生成三方session */
type TaobaoOpenlinkSessionGetAPIRequest struct {
	model.Params
	// 授权码
	_code string
}

// New
