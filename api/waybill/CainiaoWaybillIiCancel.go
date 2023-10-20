package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiCancel 商家取消获取的电子面单号
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
func CainiaoWaybillIiCancel(clt *core.SDKClient, req *waybill.CainiaoWaybillIiCancelAPIRequest, session string) (*waybill.CainiaoWaybillIiCancelAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
