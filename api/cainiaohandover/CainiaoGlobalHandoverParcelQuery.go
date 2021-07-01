package cainiaohandover

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaohandover"
)

/* 
获取交接单小包信息 
cainiao.global.handover.parcel.query

提供给ISV通过该接口查询小包信息
*/
func CainiaoGlobalHandoverParcelQuery(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIResponse, error) {
    var resp cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
