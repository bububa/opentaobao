package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
天猫好房楼盘评测同步 
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息
*/
func AlibabaAlihouseNewhomeReviewSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewSyncRequest, session string) (*alihouse.AlibabaAlihouseNewhomeReviewSyncAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeReviewSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
