package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpLabelFindconfiglist 查询可用标签id信息
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
func TaobaoUniversalbpLabelFindconfiglist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpLabelFindconfiglistAPIRequest, resp *simba.TaobaoUniversalbpLabelFindconfiglistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
