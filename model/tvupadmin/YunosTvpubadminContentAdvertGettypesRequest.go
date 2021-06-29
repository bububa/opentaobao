package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告位类型 APIRequest
yunos.tvpubadmin.content.advert.gettypes

获取广告位类型
*/
type YunosTvpubadminContentAdvertGettypesRequest struct {
    model.Params

}

func NewYunosTvpubadminContentAdvertGettypesRequest() *YunosTvpubadminContentAdvertGettypesRequest{
    return &YunosTvpubadminContentAdvertGettypesRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAdvertGettypesRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.gettypes"
}

func (r YunosTvpubadminContentAdvertGettypesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


