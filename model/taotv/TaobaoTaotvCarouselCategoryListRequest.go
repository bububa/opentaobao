package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取轮播分类列表 API请求
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
type TaobaoTaotvCarouselCategoryListRequest struct {
    model.Params
    // 设备信息
    _systemInfo   string
}

// 初始化TaobaoTaotvCarouselCategoryListRequest对象
func NewTaobaoTaotvCarouselCategoryListRequest() *TaobaoTaotvCarouselCategoryListRequest{
    return &TaobaoTaotvCarouselCategoryListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselCategoryListRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.category.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselCategoryListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 设备信息
func (r *TaobaoTaotvCarouselCategoryListRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvCarouselCategoryListRequest) GetSystemInfo() string {
    return r._systemInfo
}
