package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗状态 API请求
alibaba.scbp.showcase.status

查询橱窗状态，如总数、可用数量
*/
type AlibabaScbpShowcaseStatusRequest struct {
    model.Params
}

// 初始化AlibabaScbpShowcaseStatusRequest对象
func NewAlibabaScbpShowcaseStatusRequest() *AlibabaScbpShowcaseStatusRequest{
    return &AlibabaScbpShowcaseStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseStatusRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
