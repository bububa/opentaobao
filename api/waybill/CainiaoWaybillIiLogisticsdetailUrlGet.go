package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
电子面单物流详情授权url获取 
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url
*/
func CainiaoWaybillIiLogisticsdetailUrlGet(clt *core.SDKClient, req *waybill.CainiaoWaybillIiLogisticsdetailUrlGetRequest, session string) (*waybill.CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse, error) {
    var resp waybill.CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
