package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectTradeitem 新二租品同步
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
func AlibabaAlihouseNewhomeProjectTradeitem(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectTradeitemAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectTradeitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
