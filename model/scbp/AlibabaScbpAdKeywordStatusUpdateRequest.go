package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词启动暂停推广 API请求
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

// 初始化AlibabaScbpAdKeywordStatusUpdateRequest对象
func NewAlibabaScbpAdKeywordStatusUpdateRequest() *AlibabaScbpAdKeywordStatusUpdateRequest{
    return &AlibabaScbpAdKeywordStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdKeyword Setter
// 只能取ascci字符
func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetAdKeyword(adKeyword string) error {
    r.adKeyword = adKeyword
    r.Set("ad_keyword", adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetAdKeyword() string {
    return r.adKeyword
}
// Status Setter
// 只能去in_promotion或者stopped
func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetStatus() string {
    return r.status
}
