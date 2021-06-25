package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城市Suggest接口 APIRequest
alitrip.hotel.search.citysuggest.get

城市Suggest接口
*/
type AlitripHotelSearchCitysuggestGetRequest struct {
    model.Params

    // 用户输入的词
    keyWords   string 

    // 用户id
    userId   int64 

}

func NewAlitripHotelSearchCitysuggestGetRequest() *AlitripHotelSearchCitysuggestGetRequest{
    return &AlitripHotelSearchCitysuggestGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelSearchCitysuggestGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.search.citysuggest.get"
}

func (r AlitripHotelSearchCitysuggestGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelSearchCitysuggestGetRequest) SetKeyWords(keyWords string) error {
    r.keyWords = keyWords
    r.Set("key_words", keyWords)
    return nil
}

func (r AlitripHotelSearchCitysuggestGetRequest) GetKeyWords() string {
    return r.keyWords
}

func (r *AlitripHotelSearchCitysuggestGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripHotelSearchCitysuggestGetRequest) GetUserId() int64 {
    return r.userId
}

