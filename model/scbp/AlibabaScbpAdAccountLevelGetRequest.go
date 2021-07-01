package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广账户等级 API请求
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
type AlibabaScbpAdAccountLevelGetAPIRequest struct {
    model.Params
}

// 初始化AlibabaScbpAdAccountLevelGetAPIRequest对象
func NewAlibabaScbpAdAccountLevelGetRequest() *AlibabaScbpAdAccountLevelGetAPIRequest{
    return &AlibabaScbpAdAccountLevelGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountLevelGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.account.level.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountLevelGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
