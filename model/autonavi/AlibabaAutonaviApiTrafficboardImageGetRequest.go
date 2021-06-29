package autonavi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交通看板-栅格情报获取 API请求
alibaba.autonavi.api.trafficboard.image.get

获取指定情报板ID的二进制数据（图片）
*/
type AlibabaAutonaviApiTrafficboardImageGetRequest struct {
    model.Params
    // 设备id,  按照userid 的配置，决定是否需要
    _deviceid   string
    // 批次,终端批次，按照userid 的配置，决定是否需要
    _batch   string
    // 图片 id
    _panelid   string
    // 图像尺寸（可选）,默认尺寸为原始大小(960x600) 参数为:width  x   height   (例如:960x600),参数不正确时返回原始大小
    _size   string
    // 是否为宽高等比例（可选）,参数值 true（默认）,表示宽高等比例缩放 false:  按请求尺寸缩放
    _whscale   string
    // 城市编码
    _adcodes   string
}

// 初始化AlibabaAutonaviApiTrafficboardImageGetRequest对象
func NewAlibabaAutonaviApiTrafficboardImageGetRequest() *AlibabaAutonaviApiTrafficboardImageGetRequest{
    return &AlibabaAutonaviApiTrafficboardImageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetApiMethodName() string {
    return "alibaba.autonavi.api.trafficboard.image.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Deviceid Setter
// 设备id,  按照userid 的配置，决定是否需要
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetDeviceid(_deviceid string) error {
    r._deviceid = _deviceid
    r.Set("deviceid", _deviceid)
    return nil
}

// Deviceid Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetDeviceid() string {
    return r._deviceid
}
// Batch Setter
// 批次,终端批次，按照userid 的配置，决定是否需要
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetBatch(_batch string) error {
    r._batch = _batch
    r.Set("batch", _batch)
    return nil
}

// Batch Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetBatch() string {
    return r._batch
}
// Panelid Setter
// 图片 id
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetPanelid(_panelid string) error {
    r._panelid = _panelid
    r.Set("panelid", _panelid)
    return nil
}

// Panelid Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetPanelid() string {
    return r._panelid
}
// Size Setter
// 图像尺寸（可选）,默认尺寸为原始大小(960x600) 参数为:width  x   height   (例如:960x600),参数不正确时返回原始大小
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetSize(_size string) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetSize() string {
    return r._size
}
// Whscale Setter
// 是否为宽高等比例（可选）,参数值 true（默认）,表示宽高等比例缩放 false:  按请求尺寸缩放
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetWhscale(_whscale string) error {
    r._whscale = _whscale
    r.Set("whscale", _whscale)
    return nil
}

// Whscale Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetWhscale() string {
    return r._whscale
}
// Adcodes Setter
// 城市编码
func (r *AlibabaAutonaviApiTrafficboardImageGetRequest) SetAdcodes(_adcodes string) error {
    r._adcodes = _adcodes
    r.Set("adcodes", _adcodes)
    return nil
}

// Adcodes Getter
func (r AlibabaAutonaviApiTrafficboardImageGetRequest) GetAdcodes() string {
    return r._adcodes
}
