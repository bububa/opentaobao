package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定接口V2 API请求
taobao.usergrowth.dhh.delivery.ask

提供给媒体在曝光广告前调用
*/
type TaobaoUsergrowthDhhDeliveryAskRequest struct {
    model.Params
    // 预留json参数，与手淘团队单独沟通
    _profile   string
    // oaid的md5值， 32位小写
    _oaidMd5   string
    // idfa的md5值， 32位小写
    _idfaMd5   string
    // imei的md5值， 32位小写
    _imeiMd5   string
    // oaid的原生值
    _oaid   string
    // idfa的原生值
    _idfa   string
    // imei的原生值
    _imei   string
    // 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
    _os   string
    // 渠道标识，在大航海平台申请
    _channel   string
    // 大航海广告位，在大航海平台申请
    _advertisingSpaceId   string
}

// 初始化TaobaoUsergrowthDhhDeliveryAskRequest对象
func NewTaobaoUsergrowthDhhDeliveryAskRequest() *TaobaoUsergrowthDhhDeliveryAskRequest{
    return &TaobaoUsergrowthDhhDeliveryAskRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.dhh.delivery.ask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Profile Setter
// 预留json参数，与手淘团队单独沟通
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetProfile(_profile string) error {
    r._profile = _profile
    r.Set("profile", _profile)
    return nil
}

// Profile Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetProfile() string {
    return r._profile
}
// OaidMd5 Setter
// oaid的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaidMd5(_oaidMd5 string) error {
    r._oaidMd5 = _oaidMd5
    r.Set("oaid_md5", _oaidMd5)
    return nil
}

// OaidMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOaidMd5() string {
    return r._oaidMd5
}
// IdfaMd5 Setter
// idfa的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfaMd5(_idfaMd5 string) error {
    r._idfaMd5 = _idfaMd5
    r.Set("idfa_md5", _idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetIdfaMd5() string {
    return r._idfaMd5
}
// ImeiMd5 Setter
// imei的md5值， 32位小写
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetImeiMd5(_imeiMd5 string) error {
    r._imeiMd5 = _imeiMd5
    r.Set("imei_md5", _imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetImeiMd5() string {
    return r._imeiMd5
}
// Oaid Setter
// oaid的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaid(_oaid string) error {
    r._oaid = _oaid
    r.Set("oaid", _oaid)
    return nil
}

// Oaid Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOaid() string {
    return r._oaid
}
// Idfa Setter
// idfa的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfa(_idfa string) error {
    r._idfa = _idfa
    r.Set("idfa", _idfa)
    return nil
}

// Idfa Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetIdfa() string {
    return r._idfa
}
// Imei Setter
// imei的原生值
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetImei(_imei string) error {
    r._imei = _imei
    r.Set("imei", _imei)
    return nil
}

// Imei Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetImei() string {
    return r._imei
}
// Os Setter
// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOs(_os string) error {
    r._os = _os
    r.Set("os", _os)
    return nil
}

// Os Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOs() string {
    return r._os
}
// Channel Setter
// 渠道标识，在大航海平台申请
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetChannel() string {
    return r._channel
}
// AdvertisingSpaceId Setter
// 大航海广告位，在大航海平台申请
func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetAdvertisingSpaceId(_advertisingSpaceId string) error {
    r._advertisingSpaceId = _advertisingSpaceId
    r.Set("advertising_space_id", _advertisingSpaceId)
    return nil
}

// AdvertisingSpaceId Getter
func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetAdvertisingSpaceId() string {
    return r._advertisingSpaceId
}
