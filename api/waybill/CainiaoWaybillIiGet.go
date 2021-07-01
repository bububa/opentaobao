package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
电子面单云打印接口 
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
func CainiaoWaybillIiGet(clt *core.SDKClient, req *waybill.CainiaoWaybillIiGetAPIRequest, session string) (*waybill.CainiaoWaybillIiGetAPIResponse, error) {
    var resp waybill.CainiaoWaybillIiGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
