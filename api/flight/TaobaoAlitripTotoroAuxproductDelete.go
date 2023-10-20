package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaoalitriptotoroauxproductdelete 廉航辅营产品删除
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
func Taobaoalitriptotoroauxproductdelete(clt *core.SDKClient, req *flight.TaobaoalitriptotoroauxproductdeleteAPIRequest, session string) (*flight.TaobaoalitriptotoroauxproductdeleteAPIResponse, error) {
	var resp flight.TaobaoalitriptotoroauxproductdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
