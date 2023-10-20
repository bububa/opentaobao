package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripTravelCrsorderComplete CRS接送机商家服务完成接口
// alitrip.travel.crsorder.complete
//
// 提供给CRS接送机商家的服务完成回调接口
func AlitripTravelCrsorderComplete(clt *core.SDKClient, req *car.AlitripTravelCrsorderCompleteAPIRequest, resp *car.AlitripTravelCrsorderCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
