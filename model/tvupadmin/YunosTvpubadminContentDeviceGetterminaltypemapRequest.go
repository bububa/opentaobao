package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型表 APIRequest
yunos.tvpubadmin.content.device.getterminaltypemap

获取终端类型表
*/
type YunosTvpubadminContentDeviceGetterminaltypemapRequest struct {
    model.Params

}

func NewYunosTvpubadminContentDeviceGetterminaltypemapRequest() *YunosTvpubadminContentDeviceGetterminaltypemapRequest{
    return &YunosTvpubadminContentDeviceGetterminaltypemapRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentDeviceGetterminaltypemapRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.device.getterminaltypemap"
}

func (r YunosTvpubadminContentDeviceGetterminaltypemapRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


