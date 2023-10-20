package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamempmpprojectinitnewproject 创建新的mpproject
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
func Alibabacgamempmpprojectinitnewproject(clt *core.SDKClient, req *cloudgame.AlibabacgamempmpprojectinitnewprojectAPIRequest, session string) (*cloudgame.AlibabacgamempmpprojectinitnewprojectAPIResponse, error) {
	var resp cloudgame.AlibabacgamempmpprojectinitnewprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
