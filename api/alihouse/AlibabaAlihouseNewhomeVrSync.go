package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
VR关系数据同步 
alibaba.alihouse.newhome.vr.sync

对接易居VR关系数据迁移
*/
func AlibabaAlihouseNewhomeVrSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVrSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeVrSyncAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeVrSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
