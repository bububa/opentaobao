package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcGroupsGet 获取自定义用户分组列表
// taobao.tmc.groups.get
//
// 获取自定义用户分组列表
func TaobaoTmcGroupsGet(clt *core.SDKClient, req *tmc.TaobaoTmcGroupsGetAPIRequest, resp *tmc.TaobaoTmcGroupsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
