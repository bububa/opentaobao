package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalUnbinditem 解绑商品信息
// alibaba.alihealth.reserve.dental.unbinditem
//
// 绑定门店信息，商品信息
func AlibabaAlihealthReserveDentalUnbinditem(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalUnbinditemAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalUnbinditemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
