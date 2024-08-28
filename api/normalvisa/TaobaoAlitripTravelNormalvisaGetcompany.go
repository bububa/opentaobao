package normalvisa

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGetcompany 获取物流公司信息
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
func TaobaoAlitripTravelNormalvisaGetcompany(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
