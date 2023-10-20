package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// Taobaosubusersinfoquery 根据当前子账号登陆态，获取该子账号基本信息
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
func Taobaosubusersinfoquery(clt *core.SDKClient, req *subuser.TaobaosubusersinfoqueryAPIRequest, session string) (*subuser.TaobaosubusersinfoqueryAPIResponse, error) {
	var resp subuser.TaobaosubusersinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
