package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionMaterialCreate 小程序投放--新建投放素材
// taobao.miniapp.distribution.material.create
//
// 为可投放的小程序，增加入口的素材信息，比如图片、引导文案等等。
func TaobaoMiniappDistributionMaterialCreate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionMaterialCreateAPIRequest, resp *miniappopen.TaobaoMiniappDistributionMaterialCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
