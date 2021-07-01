package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserTopicsGetAPIRequest
获取用户开通的topic列表 API请求
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用 */
type TaobaoTmcUserTopicsGetAPIRequest struct {
	model.Params
	// 卖家nick
	_nick string
}

// New
