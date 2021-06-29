package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品 API请求
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api
*/
type AlibabaDataItemGetRequest struct {
    model.Params
    // 获取商品信息，作为客户端Weex鉴权的虚拟api
    _unNamed   string
}

// 初始化AlibabaDataItemGetRequest对象
func NewAlibabaDataItemGetRequest() *AlibabaDataItemGetRequest{
    return &AlibabaDataItemGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDataItemGetRequest) GetApiMethodName() string {
    return "alibaba.data.item.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDataItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func (r *AlibabaDataItemGetRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaDataItemGetRequest) GetUnNamed() string {
    return r._unNamed
}
