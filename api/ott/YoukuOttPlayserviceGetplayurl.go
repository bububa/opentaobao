package ott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ott"
)

/* 
获取播放串地址 
youku.ott.playservice.getplayurl

获取播放串地址服务
*/
func YoukuOttPlayserviceGetplayurl(clt *core.SDKClient, req *ott.YoukuOttPlayserviceGetplayurlAPIRequest, session string) (*ott.YoukuOttPlayserviceGetplayurlAPIResponse, error) {
    var resp ott.YoukuOttPlayserviceGetplayurlAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
