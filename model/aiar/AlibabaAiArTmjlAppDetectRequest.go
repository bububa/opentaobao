package aiar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵扫一扫入口的服务 APIRequest
alibaba.ai.ar.tmjl.app.detect

天猫精灵扫一扫入口的图像检测服务
*/
type AlibabaAiArTmjlAppDetectRequest struct {
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

func NewAlibabaAiArTmjlAppDetectRequest() *AlibabaAiArTmjlAppDetectRequest{
    return &AlibabaAiArTmjlAppDetectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAiArTmjlAppDetectRequest) GetApiMethodName() string {
    return "alibaba.ai.ar.tmjl.app.detect"
}

func (r AlibabaAiArTmjlAppDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAiArTmjlAppDetectRequest) SetImgData(imgData []*model.File) error {
    r.imgData = imgData
    r.Set("img_data", imgData)
    return nil
}

func (r AlibabaAiArTmjlAppDetectRequest) GetImgData() []*model.File {
    return r.imgData
}

func (r *AlibabaAiArTmjlAppDetectRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

func (r AlibabaAiArTmjlAppDetectRequest) GetNum() int64 {
    return r.num
}

func (r *AlibabaAiArTmjlAppDetectRequest) SetCachedTargets(cachedTargets string) error {
    r.cachedTargets = cachedTargets
    r.Set("cached_targets", cachedTargets)
    return nil
}

func (r AlibabaAiArTmjlAppDetectRequest) GetCachedTargets() string {
    return r.cachedTargets
}

func (r *AlibabaAiArTmjlAppDetectRequest) SetDeviceInfo(deviceInfo string) error {
    r.deviceInfo = deviceInfo
    r.Set("device_info", deviceInfo)
    return nil
}

func (r AlibabaAiArTmjlAppDetectRequest) GetDeviceInfo() string {
    return r.deviceInfo
}

func (r *AlibabaAiArTmjlAppDetectRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaAiArTmjlAppDetectRequest) GetVersion() string {
    return r.version
}

