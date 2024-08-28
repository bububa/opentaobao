package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSerchcrowdStateBatchUpdate 单品搜索人群修改状态
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
func TaobaoSimbaSerchcrowdStateBatchUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest, resp *simba.TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
