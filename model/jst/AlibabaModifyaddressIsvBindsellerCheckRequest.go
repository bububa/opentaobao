package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商下的商家是否开通了改地址 API请求
alibaba.modifyaddress.isv.bindseller.check

鉴权服务商和商家的绑定关系，看商家是否开通了改地址
1. 没有授权
2. 已与当前appkey签约
3. 没有签约
4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
*/
type AlibabaModifyaddressIsvBindsellerCheckRequest struct {
    model.Params
}

// 初始化AlibabaModifyaddressIsvBindsellerCheckRequest对象
func NewAlibabaModifyaddressIsvBindsellerCheckRequest() *AlibabaModifyaddressIsvBindsellerCheckRequest{
    return &AlibabaModifyaddressIsvBindsellerCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaModifyaddressIsvBindsellerCheckRequest) GetApiMethodName() string {
    return "alibaba.modifyaddress.isv.bindseller.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaModifyaddressIsvBindsellerCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
