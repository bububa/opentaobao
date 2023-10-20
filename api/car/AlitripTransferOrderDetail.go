package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Alitriptransferorderdetail 接送订单详情接口
// alitrip.transfer.order.detail
//
// 接送订单详情接口
func Alitriptransferorderdetail(clt *core.SDKClient, req *car.AlitriptransferorderdetailAPIRequest, session string) (*car.AlitriptransferorderdetailAPIResponse, error) {
	var resp car.AlitriptransferorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
