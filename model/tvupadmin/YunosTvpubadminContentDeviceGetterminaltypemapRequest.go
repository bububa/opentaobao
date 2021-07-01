package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型表 API请求
yunos.tvpubadmin.content.device.getterminaltypemap

获取终端类型表
*/
type YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest struct {
    model.Params
}

// 初始化YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest对象
func NewYunosTvpubadminContentDeviceGetterminaltypemapRequest() *YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest{
    return &YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.device.getterminaltypemap"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
