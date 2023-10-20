package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaovassubscribeget 订购关系查询
// taobao.vas.subscribe.get
//
// 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
func Taobaovassubscribeget(clt *core.SDKClient, req *servicecenter.TaobaovassubscribegetAPIRequest, session string) (*servicecenter.TaobaovassubscribegetAPIResponse, error) {
	var resp servicecenter.TaobaovassubscribegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
