package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Alitriptravelcrsdriverarrange CRS接送机商家派司机接口
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
func Alitriptravelcrsdriverarrange(clt *core.SDKClient, req *car.AlitriptravelcrsdriverarrangeAPIRequest, session string) (*car.AlitriptravelcrsdriverarrangeAPIResponse, error) {
	var resp car.AlitriptravelcrsdriverarrangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
