package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微信公众号价格变化通知 APIRequest
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

func NewAlibabaPicturesDengtaWxaccountPriceChangeRequest() *AlibabaPicturesDengtaWxaccountPriceChangeRequest{
    return &AlibabaPicturesDengtaWxaccountPriceChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.wxaccount.price.change"
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetChangeTime() string {
    return r.changeTime
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingle(single string) error {
    r.single = single
    r.Set("single", single)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingle() string {
    return r.single
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOther(other string) error {
    r.other = other
    r.Set("other", other)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOther() string {
    return r.other
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecond(second string) error {
    r.second = second
    r.Set("second", second)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecond() string {
    return r.second
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirst(first string) error {
    r.first = first
    r.Set("first", first)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirst() string {
    return r.first
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetFirstAli(firstAli string) error {
    r.firstAli = firstAli
    r.Set("first_ali", firstAli)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetFirstAli() string {
    return r.firstAli
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSecondAli(secondAli string) error {
    r.secondAli = secondAli
    r.Set("second_ali", secondAli)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSecondAli() string {
    return r.secondAli
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetSingleAli(singleAli string) error {
    r.singleAli = singleAli
    r.Set("single_ali", singleAli)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetSingleAli() string {
    return r.singleAli
}

func (r *AlibabaPicturesDengtaWxaccountPriceChangeRequest) SetOtherAli(otherAli string) error {
    r.otherAli = otherAli
    r.Set("other_ali", otherAli)
    return nil
}

func (r AlibabaPicturesDengtaWxaccountPriceChangeRequest) GetOtherAli() string {
    return r.otherAli
}

