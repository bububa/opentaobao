package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询消费抽奖配置 API请求
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

// 初始化TaobaoTvpayPromotionInfoGetRequest对象
func NewTaobaoTvpayPromotionInfoGetRequest() *TaobaoTvpayPromotionInfoGetRequest{
    return &TaobaoTvpayPromotionInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayPromotionInfoGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.promotion.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayPromotionInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayPromotionInfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayPromotionInfoGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetFrom() string {
    return r.from
}
// SubjectId Setter
// 商品id
func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubjectId(subjectId string) error {
    r.subjectId = subjectId
    r.Set("subject_id", subjectId)
    return nil
}

// SubjectId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetSubjectId() string {
    return r.subjectId
}
// ExtOrderId Setter
// 淘系订单号
func (r *TaobaoTvpayPromotionInfoGetRequest) SetExtOrderId(extOrderId string) error {
    r.extOrderId = extOrderId
    r.Set("ext_order_id", extOrderId)
    return nil
}

// ExtOrderId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetExtOrderId() string {
    return r.extOrderId
}
// IsTao Setter
// 是否淘系
func (r *TaobaoTvpayPromotionInfoGetRequest) SetIsTao(isTao bool) error {
    r.isTao = isTao
    r.Set("is_tao", isTao)
    return nil
}

// IsTao Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetIsTao() bool {
    return r.isTao
}
// Subject Setter
// 标题
func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

// Subject Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetSubject() string {
    return r.subject
}
