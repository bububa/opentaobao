package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
云打印客户端监控信息收集 
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集
*/
func CainiaoCloudprintClientinfoPut(clt *core.SDKClient, req *waybill.CainiaoCloudprintClientinfoPutRequest, session string) (*waybill.CainiaoCloudprintClientinfoPutAPIResponse, error) {
    var resp waybill.CainiaoCloudprintClientinfoPutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
