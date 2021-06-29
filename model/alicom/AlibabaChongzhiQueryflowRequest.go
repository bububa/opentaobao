package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的流量宝贝接口 API请求
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝
*/
type AlibabaChongzhiQueryflowRequest struct {
    model.Params
    // 号码
    mobile   int64
    // 来源
    clientSource   string
}

// 初始化AlibabaChongzhiQueryflowRequest对象
func NewAlibabaChongzhiQueryflowRequest() *AlibabaChongzhiQueryflowRequest{
    return &AlibabaChongzhiQueryflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaChongzhiQueryflowRequest) GetApiMethodName() string {
    return "alibaba.chongzhi.queryflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaChongzhiQueryflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 号码
func (r *AlibabaChongzhiQueryflowRequest) SetMobile(mobile int64) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaChongzhiQueryflowRequest) GetMobile() int64 {
    return r.mobile
}
// ClientSource Setter
// 来源
func (r *AlibabaChongzhiQueryflowRequest) SetClientSource(clientSource string) error {
    r.clientSource = clientSource
    r.Set("client_source", clientSource)
    return nil
}

// ClientSource Getter
func (r AlibabaChongzhiQueryflowRequest) GetClientSource() string {
    return r.clientSource
}
