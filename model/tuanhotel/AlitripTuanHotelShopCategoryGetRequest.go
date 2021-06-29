package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家店铺类目查询 API请求
alitrip.tuan.hotel.shop.category.get

查询商家店铺类目信息
*/
type AlitripTuanHotelShopCategoryGetRequest struct {
    model.Params
}

// 初始化AlitripTuanHotelShopCategoryGetRequest对象
func NewAlitripTuanHotelShopCategoryGetRequest() *AlitripTuanHotelShopCategoryGetRequest{
    return &AlitripTuanHotelShopCategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelShopCategoryGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.shop.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelShopCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
