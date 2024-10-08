package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMemberFindbrandinfolist 查询可用品牌列表
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
func TaobaoUniversalbpMemberFindbrandinfolist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpMemberFindbrandinfolistAPIRequest, resp *simba.TaobaoUniversalbpMemberFindbrandinfolistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
