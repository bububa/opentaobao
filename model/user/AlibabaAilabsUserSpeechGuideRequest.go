package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
引导语推荐接口 APIRequest
alibaba.ailabs.user.speech.guide

根据用户的语音query，返回相应的引导语推荐
*/
type AlibabaAilabsUserSpeechGuideRequest struct {
    model.Params

    // 用户query
    query   string 

}

func NewAlibabaAilabsUserSpeechGuideRequest() *AlibabaAilabsUserSpeechGuideRequest{
    return &AlibabaAilabsUserSpeechGuideRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsUserSpeechGuideRequest) GetApiMethodName() string {
    return "alibaba.ailabs.user.speech.guide"
}

func (r AlibabaAilabsUserSpeechGuideRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsUserSpeechGuideRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaAilabsUserSpeechGuideRequest) GetQuery() string {
    return r.query
}

