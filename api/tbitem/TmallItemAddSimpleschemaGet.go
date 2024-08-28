package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemAddSimpleschemaGet 天猫发布商品规则获取
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
func TmallItemAddSimpleschemaGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemAddSimpleschemaGetAPIRequest, resp *tbitem.TmallItemAddSimpleschemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
