package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiDelivery 派件通知接口
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
func CainiaoWaybillIiDelivery(clt *core.SDKClient, req *waybill.CainiaoWaybillIiDeliveryAPIRequest, session string) (*waybill.CainiaoWaybillIiDeliveryAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiDeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
