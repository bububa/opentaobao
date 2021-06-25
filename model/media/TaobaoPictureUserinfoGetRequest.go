package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询图片空间用户的信息 APIRequest
taobao.picture.userinfo.get

查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
*/
type TaobaoPictureUserinfoGetRequest struct {
    model.Params

}

func NewTaobaoPictureUserinfoGetRequest() *TaobaoPictureUserinfoGetRequest{
    return &TaobaoPictureUserinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureUserinfoGetRequest) GetApiMethodName() string {
    return "taobao.picture.userinfo.get"
}

func (r TaobaoPictureUserinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


