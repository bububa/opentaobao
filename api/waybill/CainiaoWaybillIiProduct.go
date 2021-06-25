package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
商家查询物流商产品类型接口 
cainiao.waybill.ii.product

商家可以查询物流商的产品类型和服务能力。
*/
func CainiaoWaybillIiProduct(clt *core.SDKClient, req *waybill.CainiaoWaybillIiProductRequest, session string) (*waybill.CainiaoWaybillIiProductResponse, error) {
    var resp waybill.CainiaoWaybillIiProductAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
