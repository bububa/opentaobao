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
type YunosTvpubadminContentDeviceGetterminaltypemapRequest struct {
    model.Params
}

// 初始化YunosTvpubadminContentDeviceGetterminaltypemapRequest对象
func NewYunosTvpubadminContentDeviceGetterminaltypemapRequest() *YunosTvpubadminContentDeviceGetterminaltypemapRequest{
    return &YunosTvpubadminContentDeviceGetterminaltypemapRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentDeviceGetterminaltypemapRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.device.getterminaltypemap"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentDeviceGetterminaltypemapRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
