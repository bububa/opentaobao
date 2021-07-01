package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

/* AlibabaIdleRentOrderPackage
确认揽收商品
alibaba.idle.rent.order.package

确认揽收 */
func AlibabaIdleRentOrderPackage(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderPackageAPIRequest, session string) (*idle.AlibabaIdleRentOrderPackageAPIResponse, error) {
	var resp idle.AlibabaIdleRentOrderPackageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
