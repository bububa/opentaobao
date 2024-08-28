package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeOrderstatusSync 线下门店派单以及单据相关操作接口
// taobao.jst.astrolabe.orderstatus.sync
//
// 针对ERP系统部署在门店的商家，将派单状态回流到星盘
func TaobaoJstAstrolabeOrderstatusSync(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeOrderstatusSyncAPIRequest, resp *omniorder.TaobaoJstAstrolabeOrderstatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
