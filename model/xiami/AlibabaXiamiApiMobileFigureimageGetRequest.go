package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取手机banner图 API请求
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图
*/
type AlibabaXiamiApiMobileFigureimageGetRequest struct {
    model.Params
    // 分页限制
    _limit   int64
    // 类型
    _type   string
    // 客户端版本
    _av   string
    // 设备类型
    _deviceType   string
    // 设备ID
    _deviceId   string
}

// 初始化AlibabaXiamiApiMobileFigureimageGetRequest对象
func NewAlibabaXiamiApiMobileFigureimageGetRequest() *AlibabaXiamiApiMobileFigureimageGetRequest{
    return &AlibabaXiamiApiMobileFigureimageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.mobile.figureimage.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 分页限制
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetLimit() int64 {
    return r._limit
}
// Type Setter
// 类型
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetType() string {
    return r._type
}
// Av Setter
// 客户端版本
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetAv(_av string) error {
    r._av = _av
    r.Set("av", _av)
    return nil
}

// Av Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetAv() string {
    return r._av
}
// DeviceType Setter
// 设备类型
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceType() string {
    return r._deviceType
}
// DeviceId Setter
// 设备ID
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceId() string {
    return r._deviceId
}
