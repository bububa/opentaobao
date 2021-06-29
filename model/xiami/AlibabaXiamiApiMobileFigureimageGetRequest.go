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
    limit   int64
    // 类型
    type   string
    // 客户端版本
    av   string
    // 设备类型
    deviceType   string
    // 设备ID
    deviceId   string
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
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetLimit() int64 {
    return r.limit
}
// Type Setter
// 类型
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetType() string {
    return r.type
}
// Av Setter
// 客户端版本
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetAv(av string) error {
    r.av = av
    r.Set("av", av)
    return nil
}

// Av Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetAv() string {
    return r.av
}
// DeviceType Setter
// 设备类型
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceType() string {
    return r.deviceType
}
// DeviceId Setter
// 设备ID
func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceId() string {
    return r.deviceId
}
