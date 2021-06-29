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
    userId   string
    // 搜索关键字
    keyword   string
    // 企业id
    corpId   string
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
func (r *AlitripBtripTrainCitySuggestRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlitripBtripTrainCitySuggestRequest) GetUserId() string {
    return r.userId
}
// Keyword Setter
// 搜索关键字
func (r *AlitripBtripTrainCitySuggestRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r AlitripBtripTrainCitySuggestRequest) GetKeyword() string {
    return r.keyword
}
// CorpId Setter
// 企业id
func (r *AlitripBtripTrainCitySuggestRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripTrainCitySuggestRequest) GetCorpId() string {
    return r.corpId
}
