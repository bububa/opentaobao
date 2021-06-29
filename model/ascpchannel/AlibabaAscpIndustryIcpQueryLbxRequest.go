package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icp订单号查询lbx订单号 API请求
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号
*/
type AlibabaAscpIndustryIcpQueryLbxRequest struct {
    model.Params
    // icps订单号
    icpOrderCode   string
}

// 初始化AlibabaAscpIndustryIcpQueryLbxRequest对象
func NewAlibabaAscpIndustryIcpQueryLbxRequest() *AlibabaAscpIndustryIcpQueryLbxRequest{
    return &AlibabaAscpIndustryIcpQueryLbxRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.icp.query.lbx"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IcpOrderCode Setter
// icps订单号
func (r *AlibabaAscpIndustryIcpQueryLbxRequest) SetIcpOrderCode(icpOrderCode string) error {
    r.icpOrderCode = icpOrderCode
    r.Set("icp_order_code", icpOrderCode)
    return nil
}

// IcpOrderCode Getter
func (r AlibabaAscpIndustryIcpQueryLbxRequest) GetIcpOrderCode() string {
    return r.icpOrderCode
}
