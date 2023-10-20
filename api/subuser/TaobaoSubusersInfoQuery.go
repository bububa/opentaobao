package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersInfoQuery 根据当前子账号登陆态，获取该子账号基本信息
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
func TaobaoSubusersInfoQuery(clt *core.SDKClient, req *subuser.TaobaoSubusersInfoQueryAPIRequest, session string) (*subuser.TaobaoSubusersInfoQueryAPIResponse, error) {
	var resp subuser.TaobaoSubusersInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
