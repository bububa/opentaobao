package aiar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiArTmjlAppDetectAPIRequest 天猫精灵扫一扫入口的服务 API请求
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
type AlibabaAiArTmjlAppDetectAPIRequest struct {
	model.Params
	// 原始图像数据
	_imgData *model.File
	// 最多返回的结果数，默认为1
	_num int64
	// 本地已cache的target，多个target间以|||分隔
	_cachedTargets string
	// map，描述所有设备相关信息，如设备ID，分辨率等
	_deviceInfo string
	// 版本，默认1.0
	_version string
}

// NewAlibabaAiArTmjlAppDetectRequest 初始化AlibabaAiArTmjlAppDetectAPIRequest对象
func NewAlibabaAiArTmjlAppDetectRequest() *AlibabaAiArTmjlAppDetectAPIRequest {
	return &AlibabaAiArTmjlAppDetectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.ar.tmjl.app.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImgData Setter
// 原始图像数据
func (r *AlibabaAiArTmjlAppDetectAPIRequest) SetImgData(_imgData *model.File) error {
	r._imgData = _imgData
	r.Set("img_data", _imgData)
	return nil
}

// Get ImgData Getter
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetImgData() *model.File {
	return r._imgData
}

// Set is Num Setter
// 最多返回的结果数，默认为1
func (r *AlibabaAiArTmjlAppDetectAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// Get Num Getter
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetNum() int64 {
	return r._num
}

// Set is CachedTargets Setter
// 本地已cache的target，多个target间以|||分隔
func (r *AlibabaAiArTmjlAppDetectAPIRequest) SetCachedTargets(_cachedTargets string) error {
	r._cachedTargets = _cachedTargets
	r.Set("cached_targets", _cachedTargets)
	return nil
}

// Get CachedTargets Getter
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetCachedTargets() string {
	return r._cachedTargets
}

// Set is DeviceInfo Setter
// map，描述所有设备相关信息，如设备ID，分辨率等
func (r *AlibabaAiArTmjlAppDetectAPIRequest) SetDeviceInfo(_deviceInfo string) error {
	r._deviceInfo = _deviceInfo
	r.Set("device_info", _deviceInfo)
	return nil
}

// Get DeviceInfo Getter
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetDeviceInfo() string {
	return r._deviceInfo
}

// Set is Version Setter
// 版本，默认1.0
func (r *AlibabaAiArTmjlAppDetectAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r AlibabaAiArTmjlAppDetectAPIRequest) GetVersion() string {
	return r._version
}
