package normalvisa

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGetdetail 获取单笔订单的详情
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
func TaobaoAlitripTravelNormalvisaGetdetail(ctx context.Context, clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetdetailAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaGetdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
