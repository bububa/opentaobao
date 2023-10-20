package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthdhhdeliveryaskAPIRequest 广告曝光前判定接口V2 API请求
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
type TaobaousergrowthdhhdeliveryaskAPIRequest struct {
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
	// 大航海广告位，在大航海平台申请
	_advertisingSpaceId string
	// 渠道标识，在大航海平台申请
	_channel string
	// md5加密后的caid， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本
	_caidMd5 string
}

// NewTaobaousergrowthdhhdeliveryaskRequest 初始化TaobaousergrowthdhhdeliveryaskAPIRequest对象
func NewTaobaousergrowthdhhdeliveryaskRequest() *TaobaousergrowthdhhdeliveryaskAPIRequest {
	return &TaobaousergrowthdhhdeliveryaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.dhh.delivery.ask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProfile is Profile Setter
// 预留json参数，与手淘团队单独沟通
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetProfile(_profile string) error {
	r._profile = _profile
	r.Set("profile", _profile)
	return nil
}

// GetProfile Profile Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetProfile() string {
	return r._profile
}

// SetOaidMd5 is OaidMd5 Setter
// oaid的md5值， 32位小写
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetOaidMd5(_oaidMd5 string) error {
	r._oaidMd5 = _oaidMd5
	r.Set("oaid_md5", _oaidMd5)
	return nil
}

// GetOaidMd5 OaidMd5 Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetOaidMd5() string {
	return r._oaidMd5
}

// SetIdfaMd5 is IdfaMd5 Setter
// idfa的md5值， 32位小写
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
	r._idfaMd5 = _idfaMd5
	r.Set("idfa_md5", _idfaMd5)
	return nil
}

// GetIdfaMd5 IdfaMd5 Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetIdfaMd5() string {
	return r._idfaMd5
}

// SetImeiMd5 is ImeiMd5 Setter
// imei的md5值， 32位小写
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
	r._imeiMd5 = _imeiMd5
	r.Set("imei_md5", _imeiMd5)
	return nil
}

// GetImeiMd5 ImeiMd5 Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetImeiMd5() string {
	return r._imeiMd5
}

// SetOaid is Oaid Setter
// oaid的原生值
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetOaid(_oaid string) error {
	r._oaid = _oaid
	r.Set("oaid", _oaid)
	return nil
}

// GetOaid Oaid Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetOaid() string {
	return r._oaid
}

// SetIdfa is Idfa Setter
// idfa的原生值
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetIdfa(_idfa string) error {
	r._idfa = _idfa
	r.Set("idfa", _idfa)
	return nil
}

// GetIdfa Idfa Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetIdfa() string {
	return r._idfa
}

// SetImei is Imei Setter
// imei的原生值
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetImei() string {
	return r._imei
}

// SetOs is Os Setter
// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetOs(_os string) error {
	r._os = _os
	r.Set("os", _os)
	return nil
}

// GetOs Os Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetOs() string {
	return r._os
}

// SetAdvertisingSpaceId is AdvertisingSpaceId Setter
// 大航海广告位，在大航海平台申请
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetAdvertisingSpaceId(_advertisingSpaceId string) error {
	r._advertisingSpaceId = _advertisingSpaceId
	r.Set("advertising_space_id", _advertisingSpaceId)
	return nil
}

// GetAdvertisingSpaceId AdvertisingSpaceId Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetAdvertisingSpaceId() string {
	return r._advertisingSpaceId
}

// SetChannel is Channel Setter
// 渠道标识，在大航海平台申请
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetChannel() string {
	return r._channel
}

// SetCaidMd5 is CaidMd5 Setter
// md5加密后的caid， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本
func (r *TaobaousergrowthdhhdeliveryaskAPIRequest) SetCaidMd5(_caidMd5 string) error {
	r._caidMd5 = _caidMd5
	r.Set("caid_md5", _caidMd5)
	return nil
}

// GetCaidMd5 CaidMd5 Getter
func (r TaobaousergrowthdhhdeliveryaskAPIRequest) GetCaidMd5() string {
	return r._caidMd5
}
