package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格变更接口 API请求
alibaba.mos.goods.setprice

价格变更接口，供供应商修改价格时使用。
*/
type AlibabaMosGoodsSetpriceRequest struct {
    model.Params
    // 价格变更对象列表
    priceDtoList   []PriceDto
}

// 初始化AlibabaMosGoodsSetpriceRequest对象
func NewAlibabaMosGoodsSetpriceRequest() *AlibabaMosGoodsSetpriceRequest{
    return &AlibabaMosGoodsSetpriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSetpriceRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.setprice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSetpriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PriceDtoList Setter
// 价格变更对象列表
func (r *AlibabaMosGoodsSetpriceRequest) SetPriceDtoList(priceDtoList []PriceDto) error {
    r.priceDtoList = priceDtoList
    r.Set("price_dto_list", priceDtoList)
    return nil
}

// PriceDtoList Getter
func (r AlibabaMosGoodsSetpriceRequest) GetPriceDtoList() []PriceDto {
    return r.priceDtoList
}
