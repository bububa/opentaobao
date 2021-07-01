package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthReserveDentalMarkitem
标记商品是否可预约
alibaba.alihealth.reserve.dental.markitem

标记商品是否可预约 */
func AlibabaAlihealthReserveDentalMarkitem(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIRequest, session string) (*alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
