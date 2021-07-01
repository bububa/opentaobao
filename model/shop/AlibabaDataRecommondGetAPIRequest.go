package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推荐信息 API请求
alibaba.data.recommond.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataRecommondGetAPIRequest struct {
    model.Params
    // 客户端鉴权虚拟api使用
    _unNamed   string
}

// 初始化AlibabaDataRecommondGetAPIRequest对象
func NewAlibabaDataRecommondGetRequest() *AlibabaDataRecommondGetAPIRequest{
    return &AlibabaDataRecommondGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDataRecommondGetAPIRequest) GetApiMethodName() string {
    return "alibaba.data.recommond.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDataRecommondGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabaDataRecommondGetAPIRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaDataRecommondGetAPIRequest) GetUnNamed() string {
    return r._unNamed
}
