package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的话费宝贝接口 API请求
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
type AlibabaChongzhiQueryecardsRequest struct {
    model.Params
    // 号码
    mobile   int64
    // 来源
    clientSource   string
}

// 初始化AlibabaChongzhiQueryecardsRequest对象
func NewAlibabaChongzhiQueryecardsRequest() *AlibabaChongzhiQueryecardsRequest{
    return &AlibabaChongzhiQueryecardsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaChongzhiQueryecardsRequest) GetApiMethodName() string {
    return "alibaba.chongzhi.queryecards"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaChongzhiQueryecardsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 号码
func (r *AlibabaChongzhiQueryecardsRequest) SetMobile(mobile int64) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaChongzhiQueryecardsRequest) GetMobile() int64 {
    return r.mobile
}
// ClientSource Setter
// 来源
func (r *AlibabaChongzhiQueryecardsRequest) SetClientSource(clientSource string) error {
    r.clientSource = clientSource
    r.Set("client_source", clientSource)
    return nil
}

// ClientSource Getter
func (r AlibabaChongzhiQueryecardsRequest) GetClientSource() string {
    return r.clientSource
}
