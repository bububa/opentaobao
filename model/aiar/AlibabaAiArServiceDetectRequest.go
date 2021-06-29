package aiar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ailab AR图像检索 API请求
alibaba.ai.ar.service.detect

ailab AR图像检索
*/
type AlibabaAiArServiceDetectRequest struct {
    model.Params
    // 原始图像数据
    imgData   []*model.File
    // 最多返回的结果数，默认为1
    num   int64
    // 本地已cache的target，多个target间以|||分隔
    cachedTargets   string
    // map，描述所有设备相关信息，如设备ID，分辨率等
    deviceInfo   string
    // 版本，默认1.0
    version   string
}

// 初始化AlibabaAiArServiceDetectRequest对象
func NewAlibabaAiArServiceDetectRequest() *AlibabaAiArServiceDetectRequest{
    return &AlibabaAiArServiceDetectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAiArServiceDetectRequest) GetApiMethodName() string {
    return "alibaba.ai.ar.service.detect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAiArServiceDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImgData Setter
// 原始图像数据
func (r *AlibabaAiArServiceDetectRequest) SetImgData(imgData []*model.File) error {
    r.imgData = imgData
    r.Set("img_data", imgData)
    return nil
}

// ImgData Getter
func (r AlibabaAiArServiceDetectRequest) GetImgData() []*model.File {
    return r.imgData
}
// Num Setter
// 最多返回的结果数，默认为1
func (r *AlibabaAiArServiceDetectRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r AlibabaAiArServiceDetectRequest) GetNum() int64 {
    return r.num
}
// CachedTargets Setter
// 本地已cache的target，多个target间以|||分隔
func (r *AlibabaAiArServiceDetectRequest) SetCachedTargets(cachedTargets string) error {
    r.cachedTargets = cachedTargets
    r.Set("cached_targets", cachedTargets)
    return nil
}

// CachedTargets Getter
func (r AlibabaAiArServiceDetectRequest) GetCachedTargets() string {
    return r.cachedTargets
}
// DeviceInfo Setter
// map，描述所有设备相关信息，如设备ID，分辨率等
func (r *AlibabaAiArServiceDetectRequest) SetDeviceInfo(deviceInfo string) error {
    r.deviceInfo = deviceInfo
    r.Set("device_info", deviceInfo)
    return nil
}

// DeviceInfo Getter
func (r AlibabaAiArServiceDetectRequest) GetDeviceInfo() string {
    return r.deviceInfo
}
// Version Setter
// 版本，默认1.0
func (r *AlibabaAiArServiceDetectRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r AlibabaAiArServiceDetectRequest) GetVersion() string {
    return r.version
}
