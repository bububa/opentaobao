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
type AlibabaMosGoodsSetpriceAPIRequest struct {
    model.Params
    // 价格变更对象列表
    _priceDtoList   []PriceDto
}

// 初始化AlibabaMosGoodsSetpriceAPIRequest对象
func NewAlibabaMosGoodsSetpriceRequest() *AlibabaMosGoodsSetpriceAPIRequest{
    return &AlibabaMosGoodsSetpriceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSetpriceAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.setprice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSetpriceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PriceDtoList Setter
// 价格变更对象列表
func (r *AlibabaMosGoodsSetpriceAPIRequest) SetPriceDtoList(_priceDtoList []PriceDto) error {
    r._priceDtoList = _priceDtoList
    r.Set("price_dto_list", _priceDtoList)
    return nil
}

// PriceDtoList Getter
func (r AlibabaMosGoodsSetpriceAPIRequest) GetPriceDtoList() []PriceDto {
    return r._priceDtoList
}
