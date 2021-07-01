package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道中心淘外分销品批量查询(分销商专用) API请求
alibaba.ascp.channel.distributor.product.list

此api为淘外分销的品批量查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorProductListAPIRequest struct {
    model.Params
    // 列表请求
    _productListRequest   *Productlistrequest
}

// 初始化AlibabaAscpChannelDistributorProductListAPIRequest对象
func NewAlibabaAscpChannelDistributorProductListRequest() *AlibabaAscpChannelDistributorProductListAPIRequest{
    return &AlibabaAscpChannelDistributorProductListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.product.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductListRequest Setter
// 列表请求
func (r *AlibabaAscpChannelDistributorProductListAPIRequest) SetProductListRequest(_productListRequest *Productlistrequest) error {
    r._productListRequest = _productListRequest
    r.Set("product_list_request", _productListRequest)
    return nil
}

// ProductListRequest Getter
func (r AlibabaAscpChannelDistributorProductListAPIRequest) GetProductListRequest() *Productlistrequest {
    return r._productListRequest
}
