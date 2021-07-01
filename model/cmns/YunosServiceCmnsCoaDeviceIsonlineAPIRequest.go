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
type YunosServiceCmnsCoaDeviceIsonlineAPIRequest struct {
    model.Params
    // 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
    _type   string
    // 对应的设备id值
    _value   string
}

// 初始化YunosServiceCmnsCoaDeviceIsonlineAPIRequest对象
func NewYunosServiceCmnsCoaDeviceIsonlineRequest() *YunosServiceCmnsCoaDeviceIsonlineAPIRequest{
    return &YunosServiceCmnsCoaDeviceIsonlineAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.device.isonline"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
func (r *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetType() string {
    return r._type
}
// Value Setter
// 对应的设备id值
func (r *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) SetValue(_value string) error {
    r._value = _value
    r.Set("value", _value)
    return nil
}

// Value Getter
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetValue() string {
    return r._value
}
