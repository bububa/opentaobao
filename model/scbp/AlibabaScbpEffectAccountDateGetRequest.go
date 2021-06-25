package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取最近报表生成时间 APIRequest
alibaba.scbp.effect.account.date.get

获取最近报表生成时间,格式为yyyy-MM-dd
*/
type AlibabaScbpEffectAccountDateGetRequest struct {
    model.Params

}

func NewAlibabaScbpEffectAccountDateGetRequest() *AlibabaScbpEffectAccountDateGetRequest{
    return &AlibabaScbpEffectAccountDateGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpEffectAccountDateGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.account.date.get"
}

func (r AlibabaScbpEffectAccountDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


