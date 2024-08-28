package fivee

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeInnerproductGet 国产商品资质查询
// taobao.fivee.innerproduct.get
//
// 资质共享平台，国产商品查询
func TaobaoFiveeInnerproductGet(ctx context.Context, clt *core.SDKClient, req *fivee.TaobaoFiveeInnerproductGetAPIRequest, resp *fivee.TaobaoFiveeInnerproductGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
