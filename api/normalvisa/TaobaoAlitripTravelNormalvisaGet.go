package normalvisa

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGet 获取签证记录
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
func TaobaoAlitripTravelNormalvisaGet(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
