package alime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alime"
)

// TaobaoAlimeUserTokenGet 获取用户免登录令牌
// taobao.alime.user.token.get
//
// 根据第三账号信息获取用户的免登录令牌
func TaobaoAlimeUserTokenGet(clt *core.SDKClient, req *alime.TaobaoAlimeUserTokenGetAPIRequest, resp *alime.TaobaoAlimeUserTokenGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
