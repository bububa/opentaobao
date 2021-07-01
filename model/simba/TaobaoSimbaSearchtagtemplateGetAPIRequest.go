package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSearchtagtemplateGetAPIRequest
获取搜索人群TOP用户可添加人群信息 API请求
taobao.simba.searchtagtemplate.get

获取搜索人群用户可添加人群信息 */
type TaobaoSimbaSearchtagtemplateGetAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
}

// New
