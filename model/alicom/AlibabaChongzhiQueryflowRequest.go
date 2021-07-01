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
type AlibabaChongzhiQueryflowAPIRequest struct {
    model.Params
    // 号码
    _mobile   int64
    // 来源
    _clientSource   string
}

// 初始化AlibabaChongzhiQueryflowAPIRequest对象
func NewAlibabaChongzhiQueryflowRequest() *AlibabaChongzhiQueryflowAPIRequest{
    return &AlibabaChongzhiQueryflowAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaChongzhiQueryflowAPIRequest) GetApiMethodName() string {
    return "alibaba.chongzhi.queryflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaChongzhiQueryflowAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 号码
func (r *AlibabaChongzhiQueryflowAPIRequest) SetMobile(_mobile int64) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaChongzhiQueryflowAPIRequest) GetMobile() int64 {
    return r._mobile
}
// ClientSource Setter
// 来源
func (r *AlibabaChongzhiQueryflowAPIRequest) SetClientSource(_clientSource string) error {
    r._clientSource = _clientSource
    r.Set("client_source", _clientSource)
    return nil
}

// ClientSource Getter
func (r AlibabaChongzhiQueryflowAPIRequest) GetClientSource() string {
    return r._clientSource
}
