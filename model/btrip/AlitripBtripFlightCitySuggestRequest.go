package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票城市搜索 APIRequest
alitrip.btrip.flight.city.suggest

提供机票城市搜索接口，提高OA用户对接效率
*/
type AlitripBtripFlightCitySuggestRequest struct {
    model.Params

    // 用户id
    userId   string 

    // 搜索关键字
    keyword   string 

    // 企业id
    corpId   string 

}

func NewAlitripBtripFlightCitySuggestRequest() *AlitripBtripFlightCitySuggestRequest{
    return &AlitripBtripFlightCitySuggestRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripFlightCitySuggestRequest) GetApiMethodName() string {
    return "alitrip.btrip.flight.city.suggest"
}

func (r AlitripBtripFlightCitySuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripFlightCitySuggestRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripBtripFlightCitySuggestRequest) GetUserId() string {
    return r.userId
}

func (r *AlitripBtripFlightCitySuggestRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlitripBtripFlightCitySuggestRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlitripBtripFlightCitySuggestRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripFlightCitySuggestRequest) GetCorpId() string {
    return r.corpId
}

