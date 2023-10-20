package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliiupdate 电子面单云打印更新接口
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
func Cainiaowaybilliiupdate(clt *core.SDKClient, req *waybill.CainiaowaybilliiupdateAPIRequest, session string) (*waybill.CainiaowaybilliiupdateAPIResponse, error) {
	var resp waybill.CainiaowaybilliiupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
