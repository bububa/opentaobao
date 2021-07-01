package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机扫码查询领奖状态 API请求
alibaba.alihealth.codeseller.getuseraward

贩卖机扫码查询领奖状态
*/
type AlibabaAlihealthCodesellerGetuserawardAPIRequest struct {
    model.Params
    // 追溯码
    _code   string
}

// 初始化AlibabaAlihealthCodesellerGetuserawardAPIRequest对象
func NewAlibabaAlihealthCodesellerGetuserawardRequest() *AlibabaAlihealthCodesellerGetuserawardAPIRequest{
    return &AlibabaAlihealthCodesellerGetuserawardAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthCodesellerGetuserawardAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.codeseller.getuseraward"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthCodesellerGetuserawardAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthCodesellerGetuserawardAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthCodesellerGetuserawardAPIRequest) GetCode() string {
    return r._code
}
