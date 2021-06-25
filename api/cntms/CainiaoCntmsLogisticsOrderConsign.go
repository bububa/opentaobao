package cntms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cntms"
)

/* 
菜鸟配商家仓库发货 
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货
*/
func CainiaoCntmsLogisticsOrderConsign(clt *core.SDKClient, req *cntms.CainiaoCntmsLogisticsOrderConsignRequest, session string) (*cntms.CainiaoCntmsLogisticsOrderConsignResponse, error) {
    var resp cntms.CainiaoCntmsLogisticsOrderConsignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
