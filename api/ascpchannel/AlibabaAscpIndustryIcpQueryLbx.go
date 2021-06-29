package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
icp订单号查询lbx订单号 
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号
*/
func AlibabaAscpIndustryIcpQueryLbx(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryIcpQueryLbxRequest, session string) (*ascpchannel.AlibabaAscpIndustryIcpQueryLbxAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpIndustryIcpQueryLbxAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
