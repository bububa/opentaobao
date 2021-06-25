package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取手机banner图 APIRequest
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

func NewAlibabaXiamiApiMobileFigureimageGetRequest() *AlibabaXiamiApiMobileFigureimageGetRequest{
    return &AlibabaXiamiApiMobileFigureimageGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.mobile.figureimage.get"
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetLimit() int64 {
    return r.limit
}

func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetType() string {
    return r.type
}

func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetAv(av string) error {
    r.av = av
    r.Set("av", av)
    return nil
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetAv() string {
    return r.av
}

func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *AlibabaXiamiApiMobileFigureimageGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaXiamiApiMobileFigureimageGetRequest) GetDeviceId() string {
    return r.deviceId
}

