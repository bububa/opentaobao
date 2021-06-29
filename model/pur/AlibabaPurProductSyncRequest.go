package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步产品 APIRequest
alibaba.pur.product.sync

同步产品
*/
type AlibabaPurProductSyncRequest struct {
    model.Params

    // 产品对象
    accessProductDtos   []AccessProductDto 

}

func NewAlibabaPurProductSyncRequest() *AlibabaPurProductSyncRequest{
    return &AlibabaPurProductSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurProductSyncRequest) GetApiMethodName() string {
    return "alibaba.pur.product.sync"
}

func (r AlibabaPurProductSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurProductSyncRequest) SetAccessProductDtos(accessProductDtos []AccessProductDto) error {
    r.accessProductDtos = accessProductDtos
    r.Set("access_product_dtos", accessProductDtos)
    return nil
}

func (r AlibabaPurProductSyncRequest) GetAccessProductDtos() []AccessProductDto {
    return r.accessProductDtos
}

