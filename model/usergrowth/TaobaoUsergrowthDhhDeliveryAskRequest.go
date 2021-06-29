package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定接口V2 APIRequest
taobao.usergrowth.dhh.delivery.ask

提供给媒体在曝光广告前调用
*/
type TaobaoUsergrowthDhhDeliveryAskRequest struct {
    model.Params

    // 预留json参数，与手淘团队单独沟通
    profile   string 

    // oaid的md5值， 32位小写
    oaidMd5   string 

    // idfa的md5值， 32位小写
    idfaMd5   string 

    // imei的md5值， 32位小写
    imeiMd5   string 

    // oaid的原生值
    oaid   string 

    // idfa的原生值
    idfa   string 

    // imei的原生值
    imei   string 

    // 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
    os   string 

    // 渠道标识，在大航海平台申请
    channel   string 

    // 大航海广告位，在大航海平台申请
    advertisingSpaceId   string 

}

func NewTaobaoUsergrowthDhhDeliveryAskRequest() *TaobaoUsergrowthDhhDeliveryAskRequest{
    return &TaobaoUsergrowthDhhDeliveryAskRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.dhh.delivery.ask"
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetProfile(profile string) error {
    r.profile = profile
    r.Set("profile", profile)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetProfile() string {
    return r.profile
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaidMd5(oaidMd5 string) error {
    r.oaidMd5 = oaidMd5
    r.Set("oaid_md5", oaidMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOaidMd5() string {
    return r.oaidMd5
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfaMd5(idfaMd5 string) error {
    r.idfaMd5 = idfaMd5
    r.Set("idfa_md5", idfaMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetIdfaMd5() string {
    return r.idfaMd5
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetImeiMd5(imeiMd5 string) error {
    r.imeiMd5 = imeiMd5
    r.Set("imei_md5", imeiMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetImeiMd5() string {
    return r.imeiMd5
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaid(oaid string) error {
    r.oaid = oaid
    r.Set("oaid", oaid)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOaid() string {
    return r.oaid
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfa(idfa string) error {
    r.idfa = idfa
    r.Set("idfa", idfa)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetIdfa() string {
    return r.idfa
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetImei() string {
    return r.imei
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetOs(os string) error {
    r.os = os
    r.Set("os", os)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetOs() string {
    return r.os
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoUsergrowthDhhDeliveryAskRequest) SetAdvertisingSpaceId(advertisingSpaceId string) error {
    r.advertisingSpaceId = advertisingSpaceId
    r.Set("advertising_space_id", advertisingSpaceId)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryAskRequest) GetAdvertisingSpaceId() string {
    return r.advertisingSpaceId
}

