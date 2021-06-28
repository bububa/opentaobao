package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
菜鸟工单操作回传 
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV
*/
func CainiaoCbossWorkplatformOperationReply(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformOperationReplyRequest, session string) (*logistic.CainiaoCbossWorkplatformOperationReplyAPIResponse, error) {
    var resp logistic.CainiaoCbossWorkplatformOperationReplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
