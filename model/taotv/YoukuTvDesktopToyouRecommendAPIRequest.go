package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukutvdesktoptoyourecommendAPIRequest TV桌面为你推荐接口 API请求
// youku.tv.desktop.toyou.recommend
//
// 提供为你推荐数据
type YoukutvdesktoptoyourecommendAPIRequest struct {
	model.Params
	// 用户登陆token
	_token string
	// 播控方，1:华数 7:CIBN
	_bcp string
	// 终端设备型号
	_deviceModel string
	// 终端设备mac
	_mac string
	// 终端设备uuid
	_uuid string
	// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
	_from string
	// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
	_chargeType string
	// 分辨率，sw720 sw1080
	_sw string
	// 支持媒体类型,dts,drm,dolby,h265
	_deviceMedia string
	// 请求IP
	_ip string
	// 桌面版本号
	_versionCode int64
	// 获取的最大节目数
	_maxSize int64
}

// NewYoukutvdesktoptoyourecommendRequest 初始化YoukutvdesktoptoyourecommendAPIRequest对象
func NewYoukutvdesktoptoyourecommendRequest() *YoukutvdesktoptoyourecommendAPIRequest {
	return &YoukutvdesktoptoyourecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukutvdesktoptoyourecommendAPIRequest) GetApiMethodName() string {
	return "youku.tv.desktop.toyou.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukutvdesktoptoyourecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukutvdesktoptoyourecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户登陆token
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetToken() string {
	return r._token
}

// SetBcp is Bcp Setter
// 播控方，1:华数 7:CIBN
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetBcp(_bcp string) error {
	r._bcp = _bcp
	r.Set("bcp", _bcp)
	return nil
}

// GetBcp Bcp Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetBcp() string {
	return r._bcp
}

// SetDeviceModel is DeviceModel Setter
// 终端设备型号
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetMac is Mac Setter
// 终端设备mac
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetMac() string {
	return r._mac
}

// SetUuid is Uuid Setter
// 终端设备uuid
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetUuid() string {
	return r._uuid
}

// SetFrom is From Setter
// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetFrom() string {
	return r._from
}

// SetChargeType is ChargeType Setter
// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetChargeType(_chargeType string) error {
	r._chargeType = _chargeType
	r.Set("charge_type", _chargeType)
	return nil
}

// GetChargeType ChargeType Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetChargeType() string {
	return r._chargeType
}

// SetSw is Sw Setter
// 分辨率，sw720 sw1080
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetSw(_sw string) error {
	r._sw = _sw
	r.Set("sw", _sw)
	return nil
}

// GetSw Sw Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetSw() string {
	return r._sw
}

// SetDeviceMedia is DeviceMedia Setter
// 支持媒体类型,dts,drm,dolby,h265
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetDeviceMedia(_deviceMedia string) error {
	r._deviceMedia = _deviceMedia
	r.Set("device_media", _deviceMedia)
	return nil
}

// GetDeviceMedia DeviceMedia Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetDeviceMedia() string {
	return r._deviceMedia
}

// SetIp is Ip Setter
// 请求IP
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetIp() string {
	return r._ip
}

// SetVersionCode is VersionCode Setter
// 桌面版本号
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetVersionCode(_versionCode int64) error {
	r._versionCode = _versionCode
	r.Set("version_code", _versionCode)
	return nil
}

// GetVersionCode VersionCode Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetVersionCode() int64 {
	return r._versionCode
}

// SetMaxSize is MaxSize Setter
// 获取的最大节目数
func (r *YoukutvdesktoptoyourecommendAPIRequest) SetMaxSize(_maxSize int64) error {
	r._maxSize = _maxSize
	r.Set("max_size", _maxSize)
	return nil
}

// GetMaxSize MaxSize Getter
func (r YoukutvdesktoptoyourecommendAPIRequest) GetMaxSize() int64 {
	return r._maxSize
}
