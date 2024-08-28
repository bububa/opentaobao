package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgMaterialOptional 淘宝客-推广者-物料搜索
// taobao.tbk.dg.material.optional
//
// 通用物料搜索API（导购）
func TaobaoTbkDgMaterialOptional(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgMaterialOptionalAPIRequest, resp *tbk.TaobaoTbkDgMaterialOptionalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
