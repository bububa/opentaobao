package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
查询物流公司信息 
taobao.logistics.companies.get

查询淘宝网合作的物流公司信息，用于发货接口。
*/
func TaobaoLogisticsCompaniesGet(clt *core.SDKClient, req *logistic.TaobaoLogisticsCompaniesGetRequest, session string) (*logistic.TaobaoLogisticsCompaniesGetAPIResponse, error) {
    var resp logistic.TaobaoLogisticsCompaniesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
