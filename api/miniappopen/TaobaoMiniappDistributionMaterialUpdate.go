package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionMaterialUpdate 小程序投放 --- 更新投放素材
// taobao.miniapp.distribution.material.update
//
// 修改已录入的投放素材信息。
func TaobaoMiniappDistributionMaterialUpdate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionMaterialUpdateAPIRequest, resp *miniappopen.TaobaoMiniappDistributionMaterialUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
