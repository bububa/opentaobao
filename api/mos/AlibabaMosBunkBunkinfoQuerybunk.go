package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
根据合同号查询铺位信息 
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息
*/
func AlibabaMosBunkBunkinfoQuerybunk(clt *core.SDKClient, req *mos.AlibabaMosBunkBunkinfoQuerybunkRequest, session string) (*mos.AlibabaMosBunkBunkinfoQuerybunkResponse, error) {
    var resp mos.AlibabaMosBunkBunkinfoQuerybunkAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
