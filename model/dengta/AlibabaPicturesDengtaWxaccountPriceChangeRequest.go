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
    accountId   string
    // 变更时间
    changeTime   string
    // 单图文
    single   string
    // 多图文第三条及以后
    other   string
    // 多图文第二条
    second   string
    // 多图文第一条
    first   string
    // 账号id
    id   int64
    // 多图文第一条 折后价
    firstAli   string
    // 多图文第二条 折后价
    secondAli   string
    // 单图文 折后价
    singleAli   string
    // 多图文第三条及以后 折后价
    otherAli   string
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
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetAccountId() string {
    return r.accountId
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetChangeTime() string {
    return r.changeTime
}
// Single Setter
// 单图文
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingle(single string) error {
    r.single = single
    r.Set("single", single)
    return nil
}

// Single Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingle() string {
    return r.single
}
// Other Setter
// 多图文第三条及以后
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOther(other string) error {
    r.other = other
    r.Set("other", other)
    return nil
}

// Other Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOther() string {
    return r.other
}
// Second Setter
// 多图文第二条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecond(second string) error {
    r.second = second
    r.Set("second", second)
    return nil
}

// Second Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecond() string {
    return r.second
}
// First Setter
// 多图文第一条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirst(first string) error {
    r.first = first
    r.Set("first", first)
    return nil
}

// First Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirst() string {
    return r.first
}
// Id Setter
// 账号id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetId() int64 {
    return r.id
}
// FirstAli Setter
// 多图文第一条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirstAli(firstAli string) error {
    r.firstAli = firstAli
    r.Set("first_ali", firstAli)
    return nil
}

// FirstAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirstAli() string {
    return r.firstAli
}
// SecondAli Setter
// 多图文第二条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecondAli(secondAli string) error {
    r.secondAli = secondAli
    r.Set("second_ali", secondAli)
    return nil
}

// SecondAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecondAli() string {
    return r.secondAli
}
// SingleAli Setter
// 单图文 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingleAli(singleAli string) error {
    r.singleAli = singleAli
    r.Set("single_ali", singleAli)
    return nil
}

// SingleAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingleAli() string {
    return r.singleAli
}
// OtherAli Setter
// 多图文第三条及以后 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOtherAli(otherAli string) error {
    r.otherAli = otherAli
    r.Set("other_ali", otherAli)
    return nil
}

// OtherAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOtherAli() string {
    return r.otherAli
}
