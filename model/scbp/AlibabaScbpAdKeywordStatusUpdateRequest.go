package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词启动暂停推广 APIRequest
alibaba.scbp.ad.keyword.status.update

关键词启动暂停推广
*/
type AlibabaScbpAdKeywordStatusUpdateRequest struct {
    model.Params

    // 只能取ascci字符
    adKeyword   string 

    // 只能去in_promotion或者stopped
    status   string 

}

func NewAlibabaScbpAdKeywordStatusUpdateRequest() *AlibabaScbpAdKeywordStatusUpdateRequest{
    return &AlibabaScbpAdKeywordStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.status.update"
}

func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetAdKeyword() string {
    return r.adKeyword
}

func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetStatus() string {
    return r.status
}

