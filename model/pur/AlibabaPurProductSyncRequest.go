package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步产品 API请求
alibaba.pur.product.sync

同步产品
*/
type AlibabaPurProductSyncRequest struct {
    model.Params
    // 产品对象
    _accessProductDtos   []AccessProductDTO
}

// 初始化AlibabaPurProductSyncRequest对象
func NewAlibabaPurProductSyncRequest() *AlibabaPurProductSyncRequest{
    return &AlibabaPurProductSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurProductSyncRequest) GetApiMethodName() string {
    return "alibaba.pur.product.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurProductSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccessProductDtos Setter
// 产品对象
func (r *AlibabaPurProductSyncRequest) SetAccessProductDtos(_accessProductDtos []AccessProductDTO) error {
    r._accessProductDtos = _accessProductDtos
    r.Set("access_product_dtos", _accessProductDtos)
    return nil
}

// AccessProductDtos Getter
func (r AlibabaPurProductSyncRequest) GetAccessProductDtos() []AccessProductDTO {
    return r._accessProductDtos
}
