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
type AlibabaScbpAdAccountLevelGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAdAccountLevelGetRequest对象
func NewAlibabaScbpAdAccountLevelGetRequest() *AlibabaScbpAdAccountLevelGetRequest{
    return &AlibabaScbpAdAccountLevelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountLevelGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.account.level.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountLevelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
