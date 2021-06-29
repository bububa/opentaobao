package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id查询设备是否在线 API请求
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线
*/
type YunosServiceCmnsCoaDeviceIsonlineRequest struct {
    model.Params
    // 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
    _type   string
    // 对应的设备id值
    _value   string
}

// 初始化YunosServiceCmnsCoaDeviceIsonlineRequest对象
func NewYunosServiceCmnsCoaDeviceIsonlineRequest() *YunosServiceCmnsCoaDeviceIsonlineRequest{
    return &YunosServiceCmnsCoaDeviceIsonlineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.device.isonline"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
func (r *YunosServiceCmnsCoaDeviceIsonlineRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetType() string {
    return r._type
}
// Value Setter
// 对应的设备id值
func (r *YunosServiceCmnsCoaDeviceIsonlineRequest) SetValue(_value string) error {
    r._value = _value
    r.Set("value", _value)
    return nil
}

// Value Getter
func (r YunosServiceCmnsCoaDeviceIsonlineRequest) GetValue() string {
    return r._value
}
