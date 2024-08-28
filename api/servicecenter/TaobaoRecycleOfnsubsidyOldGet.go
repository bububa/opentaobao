package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOfnsubsidyOldGet 回收单旧机款及补贴查询
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
func TaobaoRecycleOfnsubsidyOldGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoRecycleOfnsubsidyOldGetAPIRequest, resp *servicecenter.TaobaoRecycleOfnsubsidyOldGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
