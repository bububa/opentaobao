package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasSubscribeGetAPIRequest
订购关系查询 API请求
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
type TaobaoVasSubscribeGetAPIRequest struct {
	model.Params
	// 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
	_articleCode string
	// 淘宝会员名
	_nick string
}

// New
