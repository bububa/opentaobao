package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpmemberfindbrandinfolist 查询可用品牌列表
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
func Taobaouniversalbpmemberfindbrandinfolist(clt *core.SDKClient, req *simba.TaobaouniversalbpmemberfindbrandinfolistAPIRequest, session string) (*simba.TaobaouniversalbpmemberfindbrandinfolistAPIResponse, error) {
	var resp simba.TaobaouniversalbpmemberfindbrandinfolistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
