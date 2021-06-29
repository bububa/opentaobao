package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票城市搜索 API请求
alitrip.btrip.flight.city.suggest

提供机票城市搜索接口，提高OA用户对接效率
*/
type AlitripBtripFlightCitySuggestRequest struct {
    model.Params
    // 用户id
    _userId   string
    // 搜索关键字
    _keyword   string
    // 企业id
    _corpId   string
}

// 初始化AlitripBtripFlightCitySuggestRequest对象
func NewAlitripBtripFlightCitySuggestRequest() *AlitripBtripFlightCitySuggestRequest{
    return &AlitripBtripFlightCitySuggestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightCitySuggestRequest) GetApiMethodName() string {
    return "alitrip.btrip.flight.city.suggest"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightCitySuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlitripBtripFlightCitySuggestRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlitripBtripFlightCitySuggestRequest) GetUserId() string {
    return r._userId
}
// Keyword Setter
// 搜索关键字
func (r *AlitripBtripFlightCitySuggestRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlitripBtripFlightCitySuggestRequest) GetKeyword() string {
    return r._keyword
}
// CorpId Setter
// 企业id
func (r *AlitripBtripFlightCitySuggestRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripFlightCitySuggestRequest) GetCorpId() string {
    return r._corpId
}
