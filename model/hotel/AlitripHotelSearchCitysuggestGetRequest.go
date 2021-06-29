package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城市Suggest接口 API请求
alitrip.hotel.search.citysuggest.get

城市Suggest接口
*/
type AlitripHotelSearchCitysuggestGetRequest struct {
    model.Params
    // 用户输入的词
    _keyWords   string
    // 用户id
    _userId   int64
}

// 初始化AlitripHotelSearchCitysuggestGetRequest对象
func NewAlitripHotelSearchCitysuggestGetRequest() *AlitripHotelSearchCitysuggestGetRequest{
    return &AlitripHotelSearchCitysuggestGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelSearchCitysuggestGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.search.citysuggest.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelSearchCitysuggestGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeyWords Setter
// 用户输入的词
func (r *AlitripHotelSearchCitysuggestGetRequest) SetKeyWords(_keyWords string) error {
    r._keyWords = _keyWords
    r.Set("key_words", _keyWords)
    return nil
}

// KeyWords Getter
func (r AlitripHotelSearchCitysuggestGetRequest) GetKeyWords() string {
    return r._keyWords
}
// UserId Setter
// 用户id
func (r *AlitripHotelSearchCitysuggestGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlitripHotelSearchCitysuggestGetRequest) GetUserId() int64 {
    return r._userId
}
