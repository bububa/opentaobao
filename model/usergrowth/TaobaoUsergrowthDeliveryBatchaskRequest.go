package usergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放批量询问 APIRequest
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

func NewTaobaoUsergrowthDeliveryBatchaskRequest() *TaobaoUsergrowthDeliveryBatchaskRequest{
    return &TaobaoUsergrowthDeliveryBatchaskRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.batchask"
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetAdid(adid string) error {
    r.adid = adid
    r.Set("adid", adid)
    return nil
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetAdid() string {
    return r.adid
}

func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetIdfaMd5(idfaMd5 string) error {
    r.idfaMd5 = idfaMd5
    r.Set("idfa_md5", idfaMd5)
    return nil
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetIdfaMd5() string {
    return r.idfaMd5
}

func (r *TaobaoUsergrowthDeliveryBatchaskRequest) SetImeiMd5(imeiMd5 string) error {
    r.imeiMd5 = imeiMd5
    r.Set("imei_md5", imeiMd5)
    return nil
}

func (r TaobaoUsergrowthDeliveryBatchaskRequest) GetImeiMd5() string {
    return r.imeiMd5
}

