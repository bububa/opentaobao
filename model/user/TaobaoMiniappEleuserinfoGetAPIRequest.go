package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappEleuserinfoGetAPIRequest
获取饿了么用户信息 API请求
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息 */
type TaobaoMiniappEleuserinfoGetAPIRequest struct {
	model.Params
	// 怪兽业务方
	_bizScence string
}

// New
