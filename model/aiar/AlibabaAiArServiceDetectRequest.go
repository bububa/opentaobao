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
    _imgData   []*model.File
    // 最多返回的结果数，默认为1
    _num   int64
    // 本地已cache的target，多个target间以|||分隔
    _cachedTargets   string
    // map，描述所有设备相关信息，如设备ID，分辨率等
    _deviceInfo   string
    // 版本，默认1.0
    _version   string
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
func (r *AlibabaAiArServiceDetectRequest) SetImgData(_imgData []*model.File) error {
    r._imgData = _imgData
    r.Set("img_data", _imgData)
    return nil
}

// ImgData Getter
func (r AlibabaAiArServiceDetectRequest) GetImgData() []*model.File {
    return r._imgData
}
// Num Setter
// 最多返回的结果数，默认为1
func (r *AlibabaAiArServiceDetectRequest) SetNum(_num int64) error {
    r._num = _num
    r.Set("num", _num)
    return nil
}

// Num Getter
func (r AlibabaAiArServiceDetectRequest) GetNum() int64 {
    return r._num
}
// CachedTargets Setter
// 本地已cache的target，多个target间以|||分隔
func (r *AlibabaAiArServiceDetectRequest) SetCachedTargets(_cachedTargets string) error {
    r._cachedTargets = _cachedTargets
    r.Set("cached_targets", _cachedTargets)
    return nil
}

// CachedTargets Getter
func (r AlibabaAiArServiceDetectRequest) GetCachedTargets() string {
    return r._cachedTargets
}
// DeviceInfo Setter
// map，描述所有设备相关信息，如设备ID，分辨率等
func (r *AlibabaAiArServiceDetectRequest) SetDeviceInfo(_deviceInfo string) error {
    r._deviceInfo = _deviceInfo
    r.Set("device_info", _deviceInfo)
    return nil
}

// DeviceInfo Getter
func (r AlibabaAiArServiceDetectRequest) GetDeviceInfo() string {
    return r._deviceInfo
}
// Version Setter
// 版本，默认1.0
func (r *AlibabaAiArServiceDetectRequest) SetVersion(_version string) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaAiArServiceDetectRequest) GetVersion() string {
    return r._version
}
