package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripTotoroAuxproductPush 廉航辅营产品投放
// taobao.alitrip.totoro.auxproduct.push
//
// 廉航辅营产品投放接口
func TaobaoAlitripTotoroAuxproductPush(clt *core.SDKClient, req *flight.TaobaoAlitripTotoroAuxproductPushAPIRequest, resp *flight.TaobaoAlitripTotoroAuxproductPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
