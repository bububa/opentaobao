package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripTotoroAuxproductDelete 廉航辅营产品删除
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
func TaobaoAlitripTotoroAuxproductDelete(ctx context.Context, clt *core.SDKClient, req *flight.TaobaoAlitripTotoroAuxproductDeleteAPIRequest, resp *flight.TaobaoAlitripTotoroAuxproductDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
