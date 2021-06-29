package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的权益 API请求
alibaba.icbu.rfq.myequity

查询供应商权益接口
*/
type AlibabaIcbuRfqMyequityRequest struct {
    model.Params
}

// 初始化AlibabaIcbuRfqMyequityRequest对象
func NewAlibabaIcbuRfqMyequityRequest() *AlibabaIcbuRfqMyequityRequest{
    return &AlibabaIcbuRfqMyequityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqMyequityRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.myequity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqMyequityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
