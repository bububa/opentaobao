package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告位类型 API请求
yunos.tvpubadmin.content.advert.gettypes

获取广告位类型
*/
type YunosTvpubadminContentAdvertGettypesRequest struct {
    model.Params
}

// 初始化YunosTvpubadminContentAdvertGettypesRequest对象
func NewYunosTvpubadminContentAdvertGettypesRequest() *YunosTvpubadminContentAdvertGettypesRequest{
    return &YunosTvpubadminContentAdvertGettypesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertGettypesRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.gettypes"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertGettypesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
