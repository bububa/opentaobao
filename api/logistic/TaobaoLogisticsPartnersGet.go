package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
查询支持起始地到目的地范围的物流公司 
taobao.logistics.partners.get

查询物流公司信息（可以查询目的地可不可达情况）
*/
func TaobaoLogisticsPartnersGet(clt *core.SDKClient, req *logistic.TaobaoLogisticsPartnersGetRequest, session string) (*logistic.TaobaoLogisticsPartnersGetResponse, error) {
    var resp logistic.TaobaoLogisticsPartnersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
