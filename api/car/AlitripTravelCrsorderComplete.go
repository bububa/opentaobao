package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Alitriptravelcrsordercomplete CRS接送机商家服务完成接口
// alitrip.travel.crsorder.complete
//
// 提供给CRS接送机商家的服务完成回调接口
func Alitriptravelcrsordercomplete(clt *core.SDKClient, req *car.AlitriptravelcrsordercompleteAPIRequest, session string) (*car.AlitriptravelcrsordercompleteAPIResponse, error) {
	var resp car.AlitriptravelcrsordercompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
