package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliiquerybytradecode 通过订单号查询电子面单通接口
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
func Cainiaowaybilliiquerybytradecode(clt *core.SDKClient, req *waybill.CainiaowaybilliiquerybytradecodeAPIRequest, session string) (*waybill.CainiaowaybilliiquerybytradecodeAPIResponse, error) {
	var resp waybill.CainiaowaybilliiquerybytradecodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
