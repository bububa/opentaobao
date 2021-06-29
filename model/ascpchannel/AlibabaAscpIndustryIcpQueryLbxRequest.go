package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icp订单号查询lbx订单号 APIRequest
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号
*/
type AlibabaAscpIndustryIcpQueryLbxRequest struct {
    model.Params

    // icps订单号
    icpOrderCode   string 

}

func NewAlibabaAscpIndustryIcpQueryLbxRequest() *AlibabaAscpIndustryIcpQueryLbxRequest{
    return &AlibabaAscpIndustryIcpQueryLbxRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.icp.query.lbx"
}

func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpIndustryIcpQueryLbxRequest) SetIcpOrderCode(icpOrderCode string) error {
    r.icpOrderCode = icpOrderCode
    r.Set("icp_order_code", icpOrderCode)
    return nil
}

func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetIcpOrderCode() string {
    return r.icpOrderCode
}

