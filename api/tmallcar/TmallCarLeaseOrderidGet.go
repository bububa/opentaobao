package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleaseorderidget 天猫开新车查询订单id
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
func Tmallcarleaseorderidget(clt *core.SDKClient, req *tmallcar.TmallcarleaseorderidgetAPIRequest, session string) (*tmallcar.TmallcarleaseorderidgetAPIResponse, error) {
	var resp tmallcar.TmallcarleaseorderidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
