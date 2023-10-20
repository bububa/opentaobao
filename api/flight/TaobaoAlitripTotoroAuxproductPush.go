package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaoalitriptotoroauxproductpush 廉航辅营产品投放
// taobao.alitrip.totoro.auxproduct.push
//
// 廉航辅营产品投放接口
func Taobaoalitriptotoroauxproductpush(clt *core.SDKClient, req *flight.TaobaoalitriptotoroauxproductpushAPIRequest, session string) (*flight.TaobaoalitriptotoroauxproductpushAPIResponse, error) {
	var resp flight.TaobaoalitriptotoroauxproductpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
