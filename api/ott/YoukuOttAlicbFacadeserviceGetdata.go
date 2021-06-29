package ott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ott"
)

/* 
影视SDK获取设备能力值 
youku.ott.alicb.facadeservice.getdata

影视SDK获取设备能力值
*/
func YoukuOttAlicbFacadeserviceGetdata(clt *core.SDKClient, req *ott.YoukuOttAlicbFacadeserviceGetdataRequest, session string) (*ott.YoukuOttAlicbFacadeserviceGetdataAPIResponse, error) {
    var resp ott.YoukuOttAlicbFacadeserviceGetdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
