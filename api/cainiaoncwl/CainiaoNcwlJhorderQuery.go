package cainiaoncwl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaoncwl"
)

/* 
农村物流集货单查询接口 
cainiao.ncwl.jhorder.query

提供给接入商家，查询农村物流集货单
*/
func CainiaoNcwlJhorderQuery(clt *core.SDKClient, req *cainiaoncwl.CainiaoNcwlJhorderQueryRequest, session string) (*cainiaoncwl.CainiaoNcwlJhorderQueryAPIResponse, error) {
    var resp cainiaoncwl.CainiaoNcwlJhorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
