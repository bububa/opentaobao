package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthDhhDeliveryAskAPIRequest 广告曝光前判定接口V2 API请求
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
type TaobaoUsergrowthDhhDeliveryAskAPIRequest struct {
	model.Params
	// 预留json参数，与手淘团队单独沟通
	_profile string
	// oaid的md5值， 32位小写
	_oaidMd5 string
	// idfa的md5值， 32位小写
	_idfaMd5 string
	// imei的md5值， 32位小写
	_imeiMd5 string
	// oaid的原生值
	_oaid string
	// idfa的原生值
	_idfa string
	// imei的原生值
	_imei string
	// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
	_os string
	// 渠道标识，在大航海平台申请
	_channel string
	// 大航海广告位，在大航海平台申请
	_advertisingSpaceId string
}

// NewTaobaoUsergrowthDhhDeliveryAskRequest 初始化TaobaoUsergrowthDhhDeliveryAskAPIRequest对象
func NewTaobaoUsergrowthDhhDeliveryAskRequest() *TaobaoUsergrowthDhhDeliveryAskAPIRequest {
	return &TaobaoUsergrowthDhhDeliveryAskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.dhh.delivery.ask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Profile Setter
// 预留json参数，与手淘团队单独沟通
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetProfile(_profile string) error {
	r._profile = _profile
	r.Set("profile", _profile)
	return nil
}

// Get Profile Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetProfile() string {
	return r._profile
}

// Set is OaidMd5 Setter
// oaid的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetOaidMd5(_oaidMd5 string) error {
	r._oaidMd5 = _oaidMd5
	r.Set("oaid_md5", _oaidMd5)
	return nil
}

// Get OaidMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetOaidMd5() string {
	return r._oaidMd5
}

// Set is IdfaMd5 Setter
// idfa的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
	r._idfaMd5 = _idfaMd5
	r.Set("idfa_md5", _idfaMd5)
	return nil
}

// Get IdfaMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetIdfaMd5() string {
	return r._idfaMd5
}

// Set is ImeiMd5 Setter
// imei的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
	r._imeiMd5 = _imeiMd5
	r.Set("imei_md5", _imeiMd5)
	return nil
}

// Get ImeiMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetImeiMd5() string {
	return r._imeiMd5
}

// Set is Oaid Setter
// oaid的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetOaid(_oaid string) error {
	r._oaid = _oaid
	r.Set("oaid", _oaid)
	return nil
}

// Get Oaid Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetOaid() string {
	return r._oaid
}

// Set is Idfa Setter
// idfa的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetIdfa(_idfa string) error {
	r._idfa = _idfa
	r.Set("idfa", _idfa)
	return nil
}

// Get Idfa Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetIdfa() string {
	return r._idfa
}

// Set is Imei Setter
// imei的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// Get Imei Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetImei() string {
	return r._imei
}

// Set is Os Setter
// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetOs(_os string) error {
	r._os = _os
	r.Set("os", _os)
	return nil
}

// Get Os Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetOs() string {
	return r._os
}

// Set is Channel Setter
// 渠道标识，在大航海平台申请
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetChannel() string {
	return r._channel
}

// Set is AdvertisingSpaceId Setter
// 大航海广告位，在大航海平台申请
func (r *TaobaoUsergrowthDhhDeliveryAskAPIRequest) SetAdvertisingSpaceId(_advertisingSpaceId string) error {
	r._advertisingSpaceId = _advertisingSpaceId
	r.Set("advertising_space_id", _advertisingSpaceId)
	return nil
}

// Get AdvertisingSpaceId Getter
func (r TaobaoUsergrowthDhhDeliveryAskAPIRequest) GetAdvertisingSpaceId() string {
	return r._advertisingSpaceId
}
