package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Alitriptravelcrsordersearch CRS接送机订单列表搜索
// alitrip.travel.crsorder.search
//
// 提供给CRS商家搜索订单列表，仅返回订单号列表
func Alitriptravelcrsordersearch(clt *core.SDKClient, req *car.AlitriptravelcrsordersearchAPIRequest, session string) (*car.AlitriptravelcrsordersearchAPIResponse, error) {
	var resp car.AlitriptravelcrsordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
