package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpprojectInitnewproject 创建新的mpproject
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
func AlibabaCgameMpMpprojectInitnewproject(clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIRequest, session string) (*cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIResponse, error) {
	var resp cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
