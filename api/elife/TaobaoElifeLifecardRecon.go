package elife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// TaobaoElifeLifecardRecon 查询对账文件地址接口
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
func TaobaoElifeLifecardRecon(ctx context.Context, clt *core.SDKClient, req *elife.TaobaoElifeLifecardReconAPIRequest, resp *elife.TaobaoElifeLifecardReconAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
