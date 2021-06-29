package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道中心品的选品接口（淘外分销商专用） API请求
alibaba.ascp.channel.distributor.product.select

此api为淘外分销的品的选品标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorProductSelectRequest struct {
    model.Params
    // 选品请求
    _selectProductRequest   *ProductLinkRequest
}

// 初始化AlibabaAscpChannelDistributorProductSelectRequest对象
func NewAlibabaAscpChannelDistributorProductSelectRequest() *AlibabaAscpChannelDistributorProductSelectRequest{
    return &AlibabaAscpChannelDistributorProductSelectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorProductSelectRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.distributor.product.select"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorProductSelectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SelectProductRequest Setter
// 选品请求
func (r *AlibabaAscpChannelDistributorProductSelectRequest) SetSelectProductRequest(_selectProductRequest *ProductLinkRequest) error {
    r._selectProductRequest = _selectProductRequest
    r.Set("select_product_request", _selectProductRequest)
    return nil
}

// SelectProductRequest Getter
func (r AlibabaAscpChannelDistributorProductSelectRequest) GetSelectProductRequest() *ProductLinkRequest {
    return r._selectProductRequest
}
