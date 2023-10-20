package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamempmpprojectloginexistaccount 登录存在账号
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
func Alibabacgamempmpprojectloginexistaccount(clt *core.SDKClient, req *cloudgame.AlibabacgamempmpprojectloginexistaccountAPIRequest, session string) (*cloudgame.AlibabacgamempmpprojectloginexistaccountAPIResponse, error) {
	var resp cloudgame.AlibabacgamempmpprojectloginexistaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
