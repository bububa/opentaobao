package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripTravelCrsdriverArrange CRS接送机商家派司机接口
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
func AlitripTravelCrsdriverArrange(clt *core.SDKClient, req *car.AlitripTravelCrsdriverArrangeAPIRequest, session string) (*car.AlitripTravelCrsdriverArrangeAPIResponse, error) {
	var resp car.AlitripTravelCrsdriverArrangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
