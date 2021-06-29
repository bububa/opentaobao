package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道中心淘外分销品批量查询(分销商专用) APIRequest
alibaba.ascp.channel.distributor.product.list

此api为淘外分销的品批量查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorProductListRequest struct {
    model.Params

    // 列表请求
    productListRequest   *Productlistrequest 

}

func NewAlibabaAscpChannelDistributorProductListRequest() *AlibabaAscpChannelDistributorProductListRequest{
    return &AlibabaAscpChannelDistributorProductListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelDistributorProductListRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.product.list"
}

func (r AlibabaAscpChannelDistributorProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelDistributorProductListRequest) SetProductListRequest(productListRequest *Productlistrequest) error {
    r.productListRequest = productListRequest
    r.Set("product_list_request", productListRequest)
    return nil
}

func (r AlibabaAscpChannelDistributorProductListRequest) GetProductListRequest() *Productlistrequest {
    return r.productListRequest
}

