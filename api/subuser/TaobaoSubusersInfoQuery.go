package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersInfoQuery 根据当前子账号登陆态，获取该子账号基本信息
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
func TaobaoSubusersInfoQuery(clt *core.SDKClient, req *subuser.TaobaoSubusersInfoQueryAPIRequest, resp *subuser.TaobaoSubusersInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
