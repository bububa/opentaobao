package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
商圈数据同步 
alibaba.alihouse.newhome.business.sync

商圈数据同步
*/
func AlibabaAlihouseNewhomeBusinessSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeBusinessSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeBusinessSyncAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeBusinessSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
