package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票城市搜索 APIRequest
alitrip.btrip.train.city.suggest

阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
*/
type AlitripBtripTrainCitySuggestRequest struct {
    model.Params

    // 用户id
    userId   string 

    // 搜索关键字
    keyword   string 

    // 企业id
    corpId   string 

}

func NewAlitripBtripTrainCitySuggestRequest() *AlitripBtripTrainCitySuggestRequest{
    return &AlitripBtripTrainCitySuggestRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripTrainCitySuggestRequest) GetApiMethodName() string {
    return "alitrip.btrip.train.city.suggest"
}

func (r AlitripBtripTrainCitySuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripTrainCitySuggestRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripBtripTrainCitySuggestRequest) GetUserId() string {
    return r.userId
}

func (r *AlitripBtripTrainCitySuggestRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlitripBtripTrainCitySuggestRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlitripBtripTrainCitySuggestRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripTrainCitySuggestRequest) GetCorpId() string {
    return r.corpId
}

