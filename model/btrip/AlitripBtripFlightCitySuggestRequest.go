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
    userId   string
    // 搜索关键字
    keyword   string
    // 企业id
    corpId   string
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
func (r *AlitripBtripFlightCitySuggestRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlitripBtripFlightCitySuggestRequest) GetUserId() string {
    return r.userId
}
// Keyword Setter
// 搜索关键字
func (r *AlitripBtripFlightCitySuggestRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r AlitripBtripFlightCitySuggestRequest) GetKeyword() string {
    return r.keyword
}
// CorpId Setter
// 企业id
func (r *AlitripBtripFlightCitySuggestRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripFlightCitySuggestRequest) GetCorpId() string {
    return r.corpId
}
