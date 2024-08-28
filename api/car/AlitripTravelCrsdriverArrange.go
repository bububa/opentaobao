package car

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// AlitripTravelCrsdriverArrange CRS接送机商家派司机接口
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
func AlitripTravelCrsdriverArrange(ctx context.Context, clt *core.SDKClient, req *car.AlitripTravelCrsdriverArrangeAPIRequest, resp *car.AlitripTravelCrsdriverArrangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
