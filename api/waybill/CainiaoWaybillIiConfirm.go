package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliiconfirm 物流订单确认接口
// cainiao.waybill.ii.confirm
//
// 物流订单确认
func Cainiaowaybilliiconfirm(clt *core.SDKClient, req *waybill.CainiaowaybilliiconfirmAPIRequest, session string) (*waybill.CainiaowaybilliiconfirmAPIResponse, error) {
	var resp waybill.CainiaowaybilliiconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
