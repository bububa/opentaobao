package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScOptimusPromotion 淘宝客-服务商-权益物料精选
// taobao.tbk.sc.optimus.promotion
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
func TaobaoTbkScOptimusPromotion(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScOptimusPromotionAPIRequest, resp *tbk.TaobaoTbkScOptimusPromotionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
