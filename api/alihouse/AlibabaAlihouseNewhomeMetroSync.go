package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
地铁数据同步 
alibaba.alihouse.newhome.metro.sync

地铁数据同步
*/
func AlibabaAlihouseNewhomeMetroSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeMetroSyncRequest, session string) (*alihouse.AlibabaAlihouseNewhomeMetroSyncAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeMetroSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
