package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemGet 根据id查询商品
// taobao.scitem.get
//
// 根据id查询商品
func TaobaoScitemGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoScitemGetAPIRequest, resp *fenxiao.TaobaoScitemGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
