package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询渠道商api 
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api
*/
func AlibabaAlihealthTracecodesellerChannelSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
