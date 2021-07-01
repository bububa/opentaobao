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
type AlibabaScbpEffectAccountDateGetAPIRequest struct {
    model.Params
}

// 初始化AlibabaScbpEffectAccountDateGetAPIRequest对象
func NewAlibabaScbpEffectAccountDateGetRequest() *AlibabaScbpEffectAccountDateGetAPIRequest{
    return &AlibabaScbpEffectAccountDateGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectAccountDateGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.account.date.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectAccountDateGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
