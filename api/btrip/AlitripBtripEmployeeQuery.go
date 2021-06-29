package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
企业员工查询 
alitrip.btrip.employee.query

企业员工查询
*/
func AlitripBtripEmployeeQuery(clt *core.SDKClient, req *btrip.AlitripBtripEmployeeQueryRequest, session string) (*btrip.AlitripBtripEmployeeQueryAPIResponse, error) {
    var resp btrip.AlitripBtripEmployeeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
