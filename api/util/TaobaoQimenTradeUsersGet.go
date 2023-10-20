package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoQimenTradeUsersGet 获取奇门用户列表
// taobao.qimen.trade.users.get
//
// 获取已开通奇门订单服务的用户列表
func TaobaoQimenTradeUsersGet(clt *core.SDKClient, req *util.TaobaoQimenTradeUsersGetAPIRequest, resp *util.TaobaoQimenTradeUsersGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
