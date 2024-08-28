package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgOptimusPromotion 淘宝客-推广者-权益物料精选
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
func TaobaoTbkDgOptimusPromotion(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgOptimusPromotionAPIRequest, resp *tbk.TaobaoTbkDgOptimusPromotionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
