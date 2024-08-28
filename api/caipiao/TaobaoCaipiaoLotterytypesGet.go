package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoLotterytypesGet 获取可用的彩种列表
// taobao.caipiao.lotterytypes.get
//
// 获取彩票系统支持的可用于赠送的彩种列表
func TaobaoCaipiaoLotterytypesGet(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoLotterytypesGetAPIRequest, resp *caipiao.TaobaoCaipiaoLotterytypesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
