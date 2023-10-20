package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripTotoroAuxproductDelete 廉航辅营产品删除
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
func TaobaoAlitripTotoroAuxproductDelete(clt *core.SDKClient, req *flight.TaobaoAlitripTotoroAuxproductDeleteAPIRequest, resp *flight.TaobaoAlitripTotoroAuxproductDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
