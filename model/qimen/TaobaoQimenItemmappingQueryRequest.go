package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射查询接口 API请求
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
type TaobaoQimenItemmappingQueryRequest struct {
    model.Params
    // 
    _request   *Request
}

// 初始化TaobaoQimenItemmappingQueryRequest对象
func NewTaobaoQimenItemmappingQueryRequest() *TaobaoQimenItemmappingQueryRequest{
    return &TaobaoQimenItemmappingQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemmappingQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.itemmapping.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemmappingQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenItemmappingQueryRequest) SetRequest(_request *Request) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenItemmappingQueryRequest) GetRequest() *Request {
    return r._request
}
