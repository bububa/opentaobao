package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
用户测评结果回传接口 
alibaba.alihealth.alipaypfm.assess.result.sync

用户测评结果回传接口
*/
func AlibabaAlihealthAlipaypfmAssessResultSync(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthAlipaypfmAssessResultSyncAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthAlipaypfmAssessResultSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
