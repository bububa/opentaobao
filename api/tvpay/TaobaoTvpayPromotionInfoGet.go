package tvpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayPromotionInfoGet tv支付查询消费抽奖配置
// taobao.tvpay.promotion.info.get
//
// 查询消费抽奖配置
func TaobaoTvpayPromotionInfoGet(ctx context.Context, clt *core.SDKClient, req *tvpay.TaobaoTvpayPromotionInfoGetAPIRequest, resp *tvpay.TaobaoTvpayPromotionInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
