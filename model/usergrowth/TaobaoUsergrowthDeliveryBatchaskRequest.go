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
    channel   string
    // 广告id，淘宝和媒体协商
    adid   string
    // idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
    idfaMd5   string
    // imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
    imeiMd5   string
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
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetChannel() string {
    return r.channel
}
// Adid Setter
// 广告id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetAdid(adid string) error {
    r.adid = adid
    r.Set("adid", adid)
    return nil
}

// Adid Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetAdid() string {
    return r.adid
}
// IdfaMd5 Setter
// idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetIdfaMd5(idfaMd5 string) error {
    r.idfaMd5 = idfaMd5
    r.Set("idfa_md5", idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetIdfaMd5() string {
    return r.idfaMd5
}
// ImeiMd5 Setter
// imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetImeiMd5(imeiMd5 string) error {
    r.imeiMd5 = imeiMd5
    r.Set("imei_md5", imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetImeiMd5() string {
    return r.imeiMd5
}
