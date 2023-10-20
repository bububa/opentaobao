package alime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alime"
)

// TaobaoAlimeUserTokenAdvanceGet 获取用户免登录令牌v2
// taobao.alime.user.token.advance.get
//
// 根据第三账号信息获取用户的免登录令牌
func TaobaoAlimeUserTokenAdvanceGet(clt *core.SDKClient, req *alime.TaobaoAlimeUserTokenAdvanceGetAPIRequest, resp *alime.TaobaoAlimeUserTokenAdvanceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
