package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliidelivery 派件通知接口
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
func Cainiaowaybilliidelivery(clt *core.SDKClient, req *waybill.CainiaowaybilliideliveryAPIRequest, session string) (*waybill.CainiaowaybilliideliveryAPIResponse, error) {
	var resp waybill.CainiaowaybilliideliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
