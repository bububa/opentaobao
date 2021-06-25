package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取轮播分类列表 APIRequest
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
type TaobaoTaotvCarouselCategoryListRequest struct {
    model.Params

    // 设备信息
    systemInfo   string 

}

func NewTaobaoTaotvCarouselCategoryListRequest() *TaobaoTaotvCarouselCategoryListRequest{
    return &TaobaoTaotvCarouselCategoryListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvCarouselCategoryListRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.category.list"
}

func (r TaobaoTaotvCarouselCategoryListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvCarouselCategoryListRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvCarouselCategoryListRequest) GetSystemInfo() string {
    return r.systemInfo
}

