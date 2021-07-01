package shenjing

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shenjing"
)

/* 
获取神鲸活动列表 
alibaba.shenjing.core.activity.getappshowlist

获取神鲸活动列表
*/
func AlibabaShenjingCoreActivityGetappshowlist(clt *core.SDKClient, req *shenjing.AlibabaShenjingCoreActivityGetappshowlistAPIRequest, session string) (*shenjing.AlibabaShenjingCoreActivityGetappshowlistAPIResponse, error) {
    var resp shenjing.AlibabaShenjingCoreActivityGetappshowlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
