package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定批量接口V2 API请求
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2
*/
type TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest struct {
    model.Params
    // md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
    _oaidMd5   string
    // md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
    _idfaMd5   string
    // md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
    _imeiMd5   string
    // 巨浪广告位id,在巨浪平台申请
    _advertisingSpaceId   string
    // 巨浪渠道id,在巨浪平台申请
    _channel   string
}

// 初始化TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest对象
func NewTaobaoUsergrowthDhhDeliveryBatchaskRequest() *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest{
    return &TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetApiMethodName() string {
    return "taobao.usergrowth.dhh.delivery.batchask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OaidMd5 Setter
// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetOaidMd5(_oaidMd5 string) error {
    r._oaidMd5 = _oaidMd5
    r.Set("oaid_md5", _oaidMd5)
    return nil
}

// OaidMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetOaidMd5() string {
    return r._oaidMd5
}
// IdfaMd5 Setter
// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
    r._idfaMd5 = _idfaMd5
    r.Set("idfa_md5", _idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetIdfaMd5() string {
    return r._idfaMd5
}
// ImeiMd5 Setter
// md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
    r._imeiMd5 = _imeiMd5
    r.Set("imei_md5", _imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetImeiMd5() string {
    return r._imeiMd5
}
// AdvertisingSpaceId Setter
// 巨浪广告位id,在巨浪平台申请
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetAdvertisingSpaceId(_advertisingSpaceId string) error {
    r._advertisingSpaceId = _advertisingSpaceId
    r.Set("advertising_space_id", _advertisingSpaceId)
    return nil
}

// AdvertisingSpaceId Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetAdvertisingSpaceId() string {
    return r._advertisingSpaceId
}
// Channel Setter
// 巨浪渠道id,在巨浪平台申请
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetChannel() string {
    return r._channel
}
