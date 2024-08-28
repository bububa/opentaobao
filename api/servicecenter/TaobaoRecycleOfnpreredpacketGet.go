package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOfnpreredpacketGet 服务商查询前置补贴红包的最新数据
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
func TaobaoRecycleOfnpreredpacketGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoRecycleOfnpreredpacketGetAPIRequest, resp *servicecenter.TaobaoRecycleOfnpreredpacketGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
