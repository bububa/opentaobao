package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMemberFindbrandinfolist 查询可用品牌列表
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
func TaobaoUniversalbpMemberFindbrandinfolist(clt *core.SDKClient, req *simba.TaobaoUniversalbpMemberFindbrandinfolistAPIRequest, session string) (*simba.TaobaoUniversalbpMemberFindbrandinfolistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpMemberFindbrandinfolistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
