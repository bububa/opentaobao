package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取最近报表生成时间 API请求
alibaba.scbp.effect.account.date.get

获取最近报表生成时间,格式为yyyy-MM-dd
*/
type AlibabaScbpEffectAccountDateGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpEffectAccountDateGetRequest对象
func NewAlibabaScbpEffectAccountDateGetRequest() *AlibabaScbpEffectAccountDateGetRequest{
    return &AlibabaScbpEffectAccountDateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectAccountDateGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.account.date.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectAccountDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
