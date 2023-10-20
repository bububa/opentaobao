package cityretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cityretail"
)

// Taobaocityretailwmflconvertwarehouse 同城零售完美履约转仓
// taobao.cityretail.wmfl.convert.warehouse
//
// 同城零售完美履约转仓
func Taobaocityretailwmflconvertwarehouse(clt *core.SDKClient, req *cityretail.TaobaocityretailwmflconvertwarehouseAPIRequest, session string) (*cityretail.TaobaocityretailwmflconvertwarehouseAPIResponse, error) {
	var resp cityretail.TaobaocityretailwmflconvertwarehouseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
