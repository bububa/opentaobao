package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarorderquery 天猫汽车整车订单查询
// tmall.car.order.query
//
// 天猫汽车商家通过该接口查看整车订单信息
func Tmallcarorderquery(clt *core.SDKClient, req *tmallcar.TmallcarorderqueryAPIRequest, session string) (*tmallcar.TmallcarorderqueryAPIResponse, error) {
	var resp tmallcar.TmallcarorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
