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
    _adKeyword   string
    // 只能去in_promotion或者stopped
    _status   string
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
func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetAdKeyword(_adKeyword string) error {
    r._adKeyword = _adKeyword
    r.Set("ad_keyword", _adKeyword)
    return nil
}

// AdKeyword Getter
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetAdKeyword() string {
    return r._adKeyword
}
// Status Setter
// 只能去in_promotion或者stopped
func (r *AlibabaScbpAdKeywordStatusUpdateRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaScbpAdKeywordStatusUpdateRequest) GetStatus() string {
    return r._status
}
