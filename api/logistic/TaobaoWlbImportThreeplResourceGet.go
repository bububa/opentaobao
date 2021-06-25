package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
3PL直邮获取资源列表 
taobao.wlb.import.threepl.resource.get

获取3pl直邮的发货可用资源
*/
func TaobaoWlbImportThreeplResourceGet(clt *core.SDKClient, req *logistic.TaobaoWlbImportThreeplResourceGetRequest, session string) (*logistic.TaobaoWlbImportThreeplResourceGetResponse, error) {
    var resp logistic.TaobaoWlbImportThreeplResourceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
