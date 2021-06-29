package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票城市搜索 API请求
alitrip.btrip.train.city.suggest

阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
*/
type AlitripBtripTrainCitySuggestRequest struct {
    model.Params
    // 用户id
    _userId   string
    // 搜索关键字
    _keyword   string
    // 企业id
    _corpId   string
}

// 初始化AlitripBtripTrainCitySuggestRequest对象
func NewAlitripBtripTrainCitySuggestRequest() *AlitripBtripTrainCitySuggestRequest{
    return &AlitripBtripTrainCitySuggestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripTrainCitySuggestRequest) GetApiMethodName() string {
    return "alitrip.btrip.train.city.suggest"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripTrainCitySuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlitripBtripTrainCitySuggestRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlitripBtripTrainCitySuggestRequest) GetUserId() string {
    return r._userId
}
// Keyword Setter
// 搜索关键字
func (r *AlitripBtripTrainCitySuggestRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlitripBtripTrainCitySuggestRequest) GetKeyword() string {
    return r._keyword
}
// CorpId Setter
// 企业id
func (r *AlitripBtripTrainCitySuggestRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripTrainCitySuggestRequest) GetCorpId() string {
    return r._corpId
}
