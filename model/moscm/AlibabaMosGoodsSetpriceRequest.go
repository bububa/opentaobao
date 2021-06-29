package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格变更接口 APIRequest
alibaba.mos.goods.setprice

价格变更接口，供供应商修改价格时使用。
*/
type AlibabaMosGoodsSetpriceRequest struct {
    model.Params

    // 价格变更对象列表
    priceDtoList   []PriceDto 

}

func NewAlibabaMosGoodsSetpriceRequest() *AlibabaMosGoodsSetpriceRequest{
    return &AlibabaMosGoodsSetpriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosGoodsSetpriceRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.setprice"
}

func (r AlibabaMosGoodsSetpriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosGoodsSetpriceRequest) SetPriceDtoList(priceDtoList []PriceDto) error {
    r.priceDtoList = priceDtoList
    r.Set("price_dto_list", priceDtoList)
    return nil
}

func (r AlibabaMosGoodsSetpriceRequest) GetPriceDtoList() []PriceDto {
    return r.priceDtoList
}

