package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpprojectLoginexistaccount 登录存在账号
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
func AlibabaCgameMpMpprojectLoginexistaccount(clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpprojectLoginexistaccountAPIRequest, session string) (*cloudgame.AlibabaCgameMpMpprojectLoginexistaccountAPIResponse, error) {
	var resp cloudgame.AlibabaCgameMpMpprojectLoginexistaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
