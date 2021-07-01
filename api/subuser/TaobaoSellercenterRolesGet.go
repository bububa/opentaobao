package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

/* TaobaoSellercenterRolesGet
获取指定卖家的角色列表
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。 */
func TaobaoSellercenterRolesGet(clt *core.SDKClient, req *subuser.TaobaoSellercenterRolesGetAPIRequest, session string) (*subuser.TaobaoSellercenterRolesGetAPIResponse, error) {
	var resp subuser.TaobaoSellercenterRolesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
