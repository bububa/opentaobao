package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
*/
func CainiaoCbossWorkplatformBiztypeQueryall(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformBiztypeQueryallRequest, session string) (*logistic.CainiaoCbossWorkplatformBiztypeQueryallResponse, error) {
    var resp logistic.CainiaoCbossWorkplatformBiztypeQueryallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
