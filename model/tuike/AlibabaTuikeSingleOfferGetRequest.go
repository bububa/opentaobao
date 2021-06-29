package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询1688推客平台卖家推广中的商品信息 APIRequest
alibaba.tuike.single.offer.get

查询单个推客商品信息的接口
*/
type AlibabaTuikeSingleOfferGetRequest struct {
    model.Params

    // 推客id
    loginId   string 

    // 商品id
    offerId   int64 

}

func NewAlibabaTuikeSingleOfferGetRequest() *AlibabaTuikeSingleOfferGetRequest{
    return &AlibabaTuikeSingleOfferGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTuikeSingleOfferGetRequest) GetApiMethodName() string {
    return "alibaba.tuike.single.offer.get"
}

func (r AlibabaTuikeSingleOfferGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTuikeSingleOfferGetRequest) SetLoginId(loginId string) error {
    r.loginId = loginId
    r.Set("login_id", loginId)
    return nil
}

func (r AlibabaTuikeSingleOfferGetRequest) GetLoginId() string {
    return r.loginId
}

func (r *AlibabaTuikeSingleOfferGetRequest) SetOfferId(offerId int64) error {
    r.offerId = offerId
    r.Set("offer_id", offerId)
    return nil
}

func (r AlibabaTuikeSingleOfferGetRequest) GetOfferId() int64 {
    return r.offerId
}

