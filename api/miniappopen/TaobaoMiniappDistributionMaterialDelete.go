package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionMaterialDelete 小程序投放 --- 删除投放素材
// taobao.miniapp.distribution.material.delete
//
// 删除已录入的投放入口素材信息。
func TaobaoMiniappDistributionMaterialDelete(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionMaterialDeleteAPIRequest, resp *miniappopen.TaobaoMiniappDistributionMaterialDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
