package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
引导语推荐接口 API请求
alibaba.ailabs.user.speech.guide

根据用户的语音query，返回相应的引导语推荐
*/
type AlibabaAilabsUserSpeechGuideRequest struct {
    model.Params
    // 用户query
    _query   string
}

// 初始化AlibabaAilabsUserSpeechGuideRequest对象
func NewAlibabaAilabsUserSpeechGuideRequest() *AlibabaAilabsUserSpeechGuideRequest{
    return &AlibabaAilabsUserSpeechGuideRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsUserSpeechGuideRequest) GetApiMethodName() string {
    return "alibaba.ailabs.user.speech.guide"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsUserSpeechGuideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 用户query
func (r *AlibabaAilabsUserSpeechGuideRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaAilabsUserSpeechGuideRequest) GetQuery() string {
    return r._query
}
