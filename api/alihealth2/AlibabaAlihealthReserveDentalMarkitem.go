package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalMarkitem 标记商品是否可预约
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
func AlibabaAlihealthReserveDentalMarkitem(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
