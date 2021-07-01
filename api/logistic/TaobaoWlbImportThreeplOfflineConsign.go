package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
3PL直邮线下发货 
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货
*/
func TaobaoWlbImportThreeplOfflineConsign(clt *core.SDKClient, req *logistic.TaobaoWlbImportThreeplOfflineConsignAPIRequest, session string) (*logistic.TaobaoWlbImportThreeplOfflineConsignAPIResponse, error) {
    var resp logistic.TaobaoWlbImportThreeplOfflineConsignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
