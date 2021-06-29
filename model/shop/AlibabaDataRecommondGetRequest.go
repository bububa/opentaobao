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
type AlibabaDataRecommondGetRequest struct {
    model.Params
    // 客户端鉴权虚拟api使用
    unNamed   string
}

// 初始化AlibabaDataRecommondGetRequest对象
func NewAlibabaDataRecommondGetRequest() *AlibabaDataRecommondGetRequest{
    return &AlibabaDataRecommondGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDataRecommondGetRequest) GetApiMethodName() string {
    return "alibaba.data.recommond.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDataRecommondGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabaDataRecommondGetRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaDataRecommondGetRequest) GetUnNamed() string {
    return r.unNamed
}
