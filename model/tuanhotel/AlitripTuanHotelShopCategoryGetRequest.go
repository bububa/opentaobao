package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家店铺类目查询 APIRequest
alitrip.tuan.hotel.shop.category.get

查询商家店铺类目信息
*/
type AlitripTuanHotelShopCategoryGetRequest struct {
    model.Params

}

func NewAlitripTuanHotelShopCategoryGetRequest() *AlitripTuanHotelShopCategoryGetRequest{
    return &AlitripTuanHotelShopCategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTuanHotelShopCategoryGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.shop.category.get"
}

func (r AlitripTuanHotelShopCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


