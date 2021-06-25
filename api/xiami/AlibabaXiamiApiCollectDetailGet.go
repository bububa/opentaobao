package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
虾米音乐精选集详情接口 
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口
*/
func AlibabaXiamiApiCollectDetailGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiCollectDetailGetRequest, session string) (*xiami.AlibabaXiamiApiCollectDetailGetResponse, error) {
    var resp xiami.AlibabaXiamiApiCollectDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
