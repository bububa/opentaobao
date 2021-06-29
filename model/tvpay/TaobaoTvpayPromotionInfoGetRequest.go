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
    _deviceId   string
    // 来源
    _from   string
    // 商品id
    _subjectId   string
    // 淘系订单号
    _extOrderId   string
    // 是否淘系
    _isTao   bool
    // 标题
    _subject   string
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
func (r *TaobaoTvpayPromotionInfoGetRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayPromotionInfoGetRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetFrom() string {
    return r._from
}
// SubjectId Setter
// 商品id
func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubjectId(_subjectId string) error {
    r._subjectId = _subjectId
    r.Set("subject_id", _subjectId)
    return nil
}

// SubjectId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetSubjectId() string {
    return r._subjectId
}
// ExtOrderId Setter
// 淘系订单号
func (r *TaobaoTvpayPromotionInfoGetRequest) SetExtOrderId(_extOrderId string) error {
    r._extOrderId = _extOrderId
    r.Set("ext_order_id", _extOrderId)
    return nil
}

// ExtOrderId Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetExtOrderId() string {
    return r._extOrderId
}
// IsTao Setter
// 是否淘系
func (r *TaobaoTvpayPromotionInfoGetRequest) SetIsTao(_isTao bool) error {
    r._isTao = _isTao
    r.Set("is_tao", _isTao)
    return nil
}

// IsTao Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetIsTao() bool {
    return r._isTao
}
// Subject Setter
// 标题
func (r *TaobaoTvpayPromotionInfoGetRequest) SetSubject(_subject string) error {
    r._subject = _subject
    r.Set("subject", _subject)
    return nil
}

// Subject Getter
func (r TaobaoTvpayPromotionInfoGetRequest) GetSubject() string {
    return r._subject
}
