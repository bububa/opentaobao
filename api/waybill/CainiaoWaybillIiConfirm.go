package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiConfirm 物流订单确认接口
// cainiao.waybill.ii.confirm
//
// 物流订单确认
func CainiaoWaybillIiConfirm(clt *core.SDKClient, req *waybill.CainiaoWaybillIiConfirmAPIRequest, session string) (*waybill.CainiaoWaybillIiConfirmAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
