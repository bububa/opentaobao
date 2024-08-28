package normalvisa

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaStoreuser 代填办理人信息
// taobao.alitrip.travel.normalvisa.storeuser
//
// 卖家代填买家填写办理人信息
func TaobaoAlitripTravelNormalvisaStoreuser(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaStoreuserAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaStoreuserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
