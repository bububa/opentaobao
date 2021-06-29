package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微信公众号价格变化通知 API请求
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口
*/
type AlibabaPicturesDengtaWxaccountPriceChangeRequest struct {
    model.Params
    // id
    _accountId   string
    // 变更时间
    _changeTime   string
    // 单图文
    _single   string
    // 多图文第三条及以后
    _other   string
    // 多图文第二条
    _second   string
    // 多图文第一条
    _first   string
    // 账号id
    _id   int64
    // 多图文第一条 折后价
    _firstAli   string
    // 多图文第二条 折后价
    _secondAli   string
    // 单图文 折后价
    _singleAli   string
    // 多图文第三条及以后 折后价
    _otherAli   string
}

// 初始化AlibabaPicturesDengtaWxaccountPriceChangeRequest对象
func NewAlibabaPicturesDengtaWxaccountPriceChangeRequest() *AlibabaPicturesDengtaWxaccountPriceChangeRequest{
    return &AlibabaPicturesDengtaWxaccountPriceChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.wxaccount.price.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetAccountId() string {
    return r._accountId
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetChangeTime(_changeTime string) error {
    r._changeTime = _changeTime
    r.Set("change_time", _changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetChangeTime() string {
    return r._changeTime
}
// Single Setter
// 单图文
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingle(_single string) error {
    r._single = _single
    r.Set("single", _single)
    return nil
}

// Single Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingle() string {
    return r._single
}
// Other Setter
// 多图文第三条及以后
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOther(_other string) error {
    r._other = _other
    r.Set("other", _other)
    return nil
}

// Other Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOther() string {
    return r._other
}
// Second Setter
// 多图文第二条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecond(_second string) error {
    r._second = _second
    r.Set("second", _second)
    return nil
}

// Second Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecond() string {
    return r._second
}
// First Setter
// 多图文第一条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirst(_first string) error {
    r._first = _first
    r.Set("first", _first)
    return nil
}

// First Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirst() string {
    return r._first
}
// Id Setter
// 账号id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetId() int64 {
    return r._id
}
// FirstAli Setter
// 多图文第一条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirstAli(_firstAli string) error {
    r._firstAli = _firstAli
    r.Set("first_ali", _firstAli)
    return nil
}

// FirstAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirstAli() string {
    return r._firstAli
}
// SecondAli Setter
// 多图文第二条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecondAli(_secondAli string) error {
    r._secondAli = _secondAli
    r.Set("second_ali", _secondAli)
    return nil
}

// SecondAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecondAli() string {
    return r._secondAli
}
// SingleAli Setter
// 单图文 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingleAli(_singleAli string) error {
    r._singleAli = _singleAli
    r.Set("single_ali", _singleAli)
    return nil
}

// SingleAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingleAli() string {
    return r._singleAli
}
// OtherAli Setter
// 多图文第三条及以后 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOtherAli(_otherAli string) error {
    r._otherAli = _otherAli
    r.Set("other_ali", _otherAli)
    return nil
}

// OtherAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOtherAli() string {
    return r._otherAli
}
