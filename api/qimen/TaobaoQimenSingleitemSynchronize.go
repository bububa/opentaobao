package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenSingleitemSynchronize 商品同步接口
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
func TaobaoQimenSingleitemSynchronize(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenSingleitemSynchronizeAPIRequest, resp *qimen.TaobaoQimenSingleitemSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
