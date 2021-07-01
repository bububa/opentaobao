package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
环线数据同步 
alibaba.alihouse.newhome.line.sync

环线数据同步
*/
func AlibabaAlihouseNewhomeLineSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLineSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeLineSyncAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeLineSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
