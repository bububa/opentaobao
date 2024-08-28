package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoGoodsInfoInput 录入参加送彩票商品信息
// taobao.caipiao.goods.info.input
//
// 录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
func TaobaoCaipiaoGoodsInfoInput(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoGoodsInfoInputAPIRequest, resp *caipiao.TaobaoCaipiaoGoodsInfoInputAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
