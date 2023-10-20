package aiar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiArOpenPlatformDetectAPIRequest AR开发者平台marker图片检测服务 API请求
// alibaba.ai.ar.open.platform.detect
//
// AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用
type AlibabaAiArOpenPlatformDetectAPIRequest struct {
	model.Params
	// 本地已cache的target，多个target间以|||分隔
	_cachedTargets string
	// map，描述所有设备相关信息，如设备ID，分辨率等
	_deviceInfo string
	// 版本，默认1.0
	_version string
	// 原始图像数据
	_imgData *model.File
	// 最多返回的结果数，默认为1
	_num int64
}

// NewAlibabaAiArOpenPlatformDetectRequest 初始化AlibabaAiArOpenPlatformDetectAPIRequest对象
func NewAlibabaAiArOpenPlatformDetectRequest() *AlibabaAiArOpenPlatformDetectAPIRequest {
	return &AlibabaAiArOpenPlatformDetectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.ar.open.platform.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCachedTargets is CachedTargets Setter
// 本地已cache的target，多个target间以|||分隔
func (r *AlibabaAiArOpenPlatformDetectAPIRequest) SetCachedTargets(_cachedTargets string) error {
	r._cachedTargets = _cachedTargets
	r.Set("cached_targets", _cachedTargets)
	return nil
}

// GetCachedTargets CachedTargets Getter
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetCachedTargets() string {
	return r._cachedTargets
}

// SetDeviceInfo is DeviceInfo Setter
// map，描述所有设备相关信息，如设备ID，分辨率等
func (r *AlibabaAiArOpenPlatformDetectAPIRequest) SetDeviceInfo(_deviceInfo string) error {
	r._deviceInfo = _deviceInfo
	r.Set("device_info", _deviceInfo)
	return nil
}

// GetDeviceInfo DeviceInfo Getter
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetDeviceInfo() string {
	return r._deviceInfo
}

// SetVersion is Version Setter
// 版本，默认1.0
func (r *AlibabaAiArOpenPlatformDetectAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetVersion() string {
	return r._version
}

// SetImgData is ImgData Setter
// 原始图像数据
func (r *AlibabaAiArOpenPlatformDetectAPIRequest) SetImgData(_imgData *model.File) error {
	r._imgData = _imgData
	r.Set("img_data", _imgData)
	return nil
}

// GetImgData ImgData Getter
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetImgData() *model.File {
	return r._imgData
}

// SetNum is Num Setter
// 最多返回的结果数，默认为1
func (r *AlibabaAiArOpenPlatformDetectAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r AlibabaAiArOpenPlatformDetectAPIRequest) GetNum() int64 {
	return r._num
}
