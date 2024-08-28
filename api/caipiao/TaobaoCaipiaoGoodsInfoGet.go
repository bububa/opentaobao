package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoGoodsInfoGet 根据卖家id与appkey获取商品信息
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
func TaobaoCaipiaoGoodsInfoGet(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoGoodsInfoGetAPIRequest, resp *caipiao.TaobaoCaipiaoGoodsInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
