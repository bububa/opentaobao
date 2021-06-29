package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 API请求
yunos.tvpubadmin.device.tvid

通过UUID查询终端信息
*/
type YunosTvpubadminDeviceTvidRequest struct {
    model.Params
    // 设备的UUID
    _uuid   string
}

// 初始化YunosTvpubadminDeviceTvidRequest对象
func NewYunosTvpubadminDeviceTvidRequest() *YunosTvpubadminDeviceTvidRequest{
    return &YunosTvpubadminDeviceTvidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceTvidRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.tvid"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceTvidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备的UUID
func (r *YunosTvpubadminDeviceTvidRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YunosTvpubadminDeviceTvidRequest) GetUuid() string {
    return r._uuid
}
