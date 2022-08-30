package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// AlibabaAlscGrowthInteractiveMiniGameIndexQuery 小游戏首页查询接口
// alibaba.alsc.growth.interactive.mini.game.index.query
//
// 小游戏首页查询接口
func AlibabaAlscGrowthInteractiveMiniGameIndexQuery(clt *core.SDKClient, req *usergrowth2.AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest, session string) (*usergrowth2.AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponse, error) {
	var resp usergrowth2.AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
