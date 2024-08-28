package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractCoinBuyerAdd 平台向买家发放淘金币
// alibaba.interact.coin.buyer.add
//
// 手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。ISV调用该接口向买家发放平台淘金币，需要优惠平台运营审核ISV资质。
func AlibabaInteractCoinBuyerAdd(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractCoinBuyerAddAPIRequest, resp *interact.AlibabaInteractCoinBuyerAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
