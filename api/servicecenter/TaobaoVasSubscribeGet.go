package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

/* TaobaoVasSubscribeGet
订购关系查询
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
func TaobaoVasSubscribeGet(clt *core.SDKClient, req *servicecenter.TaobaoVasSubscribeGetAPIRequest, session string) (*servicecenter.TaobaoVasSubscribeGetAPIResponse, error) {
	var resp servicecenter.TaobaoVasSubscribeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
