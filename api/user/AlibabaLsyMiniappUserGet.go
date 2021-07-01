package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

/* AlibabaLsyMiniappUserGet
零售云小程序获取登录用户信息
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息 */
func AlibabaLsyMiniappUserGet(clt *core.SDKClient, req *user.AlibabaLsyMiniappUserGetAPIRequest, session string) (*user.AlibabaLsyMiniappUserGetAPIResponse, error) {
	var resp user.AlibabaLsyMiniappUserGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
