package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗状态 APIRequest
alibaba.scbp.showcase.status

查询橱窗状态，如总数、可用数量
*/
type AlibabaScbpShowcaseStatusRequest struct {
    model.Params

}

func NewAlibabaScbpShowcaseStatusRequest() *AlibabaScbpShowcaseStatusRequest{
    return &AlibabaScbpShowcaseStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpShowcaseStatusRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.status"
}

func (r AlibabaScbpShowcaseStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


