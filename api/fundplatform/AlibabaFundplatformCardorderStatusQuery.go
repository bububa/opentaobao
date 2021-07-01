package fundplatform

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fundplatform"
)

/* 
查询制卡商制卡进度 
alibaba.fundplatform.cardorder.status.query

当通知制卡商进行制卡后，其制卡流程会比较长，若长时间未反馈当前制卡进展，则需要使用该接口来向制卡商发起进度查询。
*/
func AlibabaFundplatformCardorderStatusQuery(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderStatusQueryAPIRequest, session string) (*fundplatform.AlibabaFundplatformCardorderStatusQueryAPIResponse, error) {
    var resp fundplatform.AlibabaFundplatformCardorderStatusQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
