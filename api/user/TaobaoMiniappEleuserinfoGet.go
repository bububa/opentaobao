package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

/* TaobaoMiniappEleuserinfoGet
获取饿了么用户信息
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息 */
func TaobaoMiniappEleuserinfoGet(clt *core.SDKClient, req *user.TaobaoMiniappEleuserinfoGetAPIRequest, session string) (*user.TaobaoMiniappEleuserinfoGetAPIResponse, error) {
	var resp user.TaobaoMiniappEleuserinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
