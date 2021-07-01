package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappUserPhoneGetAPIRequest
获取当前授权用户手机号码 API请求
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码 */
type TaobaoMiniappUserPhoneGetAPIRequest struct {
	model.Params
}

// New
