package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词推广账户是否欠款 API请求
alibaba.scbp.account.isarrears.get

查询关键词推广账户是否欠款
*/
type AlibabaScbpAccountIsarrearsGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAccountIsarrearsGetRequest对象
func NewAlibabaScbpAccountIsarrearsGetRequest() *AlibabaScbpAccountIsarrearsGetRequest{
    return &AlibabaScbpAccountIsarrearsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountIsarrearsGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.isarrears.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountIsarrearsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
