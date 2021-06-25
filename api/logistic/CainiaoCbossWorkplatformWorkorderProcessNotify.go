package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
菜鸟工单系统的工单进度下发 
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI）
*/
func CainiaoCbossWorkplatformWorkorderProcessNotify(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformWorkorderProcessNotifyRequest, session string) (*logistic.CainiaoCbossWorkplatformWorkorderProcessNotifyResponse, error) {
    var resp logistic.CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
