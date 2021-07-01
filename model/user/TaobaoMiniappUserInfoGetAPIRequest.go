package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappUserInfoGetAPIRequest
用户开放信息获取 API请求
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接 */
type TaobaoMiniappUserInfoGetAPIRequest struct {
	model.Params
}

// New
