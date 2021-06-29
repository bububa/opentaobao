package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放批量询问 API请求
taobao.usergrowth.delivery.batchask

提供给媒体在曝光广告前调用， 返回是否曝光以及报价
*/
type TaobaoUsergrowthDeliveryBatchaskRequest struct {
    model.Params
    // 渠道标识，向淘宝技术申请
    _channel   string
    // 广告id，淘宝和媒体协商
    _adid   string
    // idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
    _idfaMd5   string
    // imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
    _imeiMd5   string
}

// 初始化TaobaoUsergrowthDeliveryBatchaskRequest对象
func NewTaobaoUsergrowthDeliveryBatchaskRequest() *TaobaoUsergrowthDeliveryBatchaskRequest{
    return &TaobaoUsergrowthDeliveryBatchaskRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.batchask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Channel Setter
// 渠道标识，向淘宝技术申请
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetChannel() string {
    return r._channel
}
// Adid Setter
// 广告id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetAdid(_adid string) error {
    r._adid = _adid
    r.Set("adid", _adid)
    return nil
}

// Adid Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetAdid() string {
    return r._adid
}
// IdfaMd5 Setter
// idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetIdfaMd5(_idfaMd5 string) error {
    r._idfaMd5 = _idfaMd5
    r.Set("idfa_md5", _idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetIdfaMd5() string {
    return r._idfaMd5
}
// ImeiMd5 Setter
// imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetImeiMd5(_imeiMd5 string) error {
    r._imeiMd5 = _imeiMd5
    r.Set("imei_md5", _imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetImeiMd5() string {
    return r._imeiMd5
}
