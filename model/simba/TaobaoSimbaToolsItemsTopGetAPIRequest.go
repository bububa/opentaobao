package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaToolsItemsTopGetAPIRequest
取得一个关键词的推广组排名列表 API请求
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表 */
type TaobaoSimbaToolsItemsTopGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词
	_keyword string
	// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
	_ip string
}

// New
