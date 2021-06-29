package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道中心品的选品接口（淘外分销商专用） APIRequest
alibaba.ascp.channel.distributor.product.select

此api为淘外分销的品的选品标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorProductSelectRequest struct {
    model.Params

    // 选品请求
    selectProductRequest   *ProductLinkRequest 

}

func NewAlibabaAscpChannelDistributorProductSelectRequest() *AlibabaAscpChannelDistributorProductSelectRequest{
    return &AlibabaAscpChannelDistributorProductSelectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelDistributorProductSelectRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.product.select"
}

func (r AlibabaAscpChannelDistributorProductSelectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelDistributorProductSelectRequest) SetSelectProductRequest(selectProductRequest *ProductLinkRequest) error {
    r.selectProductRequest = selectProductRequest
    r.Set("select_product_request", selectProductRequest)
    return nil
}

func (r AlibabaAscpChannelDistributorProductSelectRequest) GetSelectProductRequest() *ProductLinkRequest {
    return r.selectProductRequest
}

