package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询消费抽奖配置 APIRequest
taobao.tvpay.promotion.info.get

查询消费抽奖配置
*/
type TaobaoTvpayPromotionInfoGetRequest struct {
    model.Params

    // 设备id
    deviceId   string 

    // 来源
    from   string 

    // 商品id
    subjectId   string 

    // 淘系订单号
    extOrderId   string 

    // 是否淘系
    isTao   bool 

    // 标题
    subject   string 

}

func NewTaobaoTvpayPromotionInfoGetRequest() *TaobaoTvpayPromotionInfoGetRequest{
    return &TaobaoTvpayPromotionInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.promotion.info.get"
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayPromotionInfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayPromotionInfoGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubjectId(subjectId string) error {
    r.subjectId = subjectId
    r.Set("subject_id", subjectId)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetSubjectId() string {
    return r.subjectId
}

func (r *TaobaoTvpayPromotionInfoGetRequest) SetExtOrderId(extOrderId string) error {
    r.extOrderId = extOrderId
    r.Set("ext_order_id", extOrderId)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetExtOrderId() string {
    return r.extOrderId
}

func (r *TaobaoTvpayPromotionInfoGetRequest) SetIsTao(isTao bool) error {
    r.isTao = isTao
    r.Set("is_tao", isTao)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetIsTao() bool {
    return r.isTao
}

func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

func (r TaobaoTvpayPromotionInfoGetRequest) GetSubject() string {
    return r.subject
}

