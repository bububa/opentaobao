package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取终端类型表 
yunos.tvpubadmin.content.device.getterminaltypemap

获取终端类型表
*/
func YunosTvpubadminContentDeviceGetterminaltypemap(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentDeviceGetterminaltypemapRequest, session string) (*tvupadmin.YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
