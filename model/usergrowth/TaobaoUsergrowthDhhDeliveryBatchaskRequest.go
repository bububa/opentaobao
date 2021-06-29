package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定批量接口V2 APIRequest
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2
*/
type TaobaoUsergrowthDhhDeliveryBatchaskRequest struct {
    model.Params

    // md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
    oaidMd5   string 

    // md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
    idfaMd5   string 

    // md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
    imeiMd5   string 

    // 巨浪广告位id,在巨浪平台申请
    advertisingSpaceId   string 

    // 巨浪渠道id,在巨浪平台申请
    channel   string 

}

func NewTaobaoUsergrowthDhhDeliveryBatchaskRequest() *TaobaoUsergrowthDhhDeliveryBatchaskRequest{
    return &TaobaoUsergrowthDhhDeliveryBatchaskRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.dhh.delivery.batchask"
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetOaidMd5(oaidMd5 string) error {
    r.oaidMd5 = oaidMd5
    r.Set("oaid_md5", oaidMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetOaidMd5() string {
    return r.oaidMd5
}

func (r *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetIdfaMd5(idfaMd5 string) error {
    r.idfaMd5 = idfaMd5
    r.Set("idfa_md5", idfaMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetIdfaMd5() string {
    return r.idfaMd5
}

func (r *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetImeiMd5(imeiMd5 string) error {
    r.imeiMd5 = imeiMd5
    r.Set("imei_md5", imeiMd5)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetImeiMd5() string {
    return r.imeiMd5
}

func (r *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetAdvertisingSpaceId(advertisingSpaceId string) error {
    r.advertisingSpaceId = advertisingSpaceId
    r.Set("advertising_space_id", advertisingSpaceId)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetAdvertisingSpaceId() string {
    return r.advertisingSpaceId
}

func (r *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoUsergrowthDhhDeliveryBatchaskRequest) GetChannel() string {
    return r.channel
}

