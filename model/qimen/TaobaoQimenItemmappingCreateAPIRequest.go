package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射接口 API请求
taobao.qimen.itemmapping.create

前后端商品映射
*/
type TaobaoQimenItemmappingCreateAPIRequest struct {
    model.Params
    // 
    _request   *RequestDo
}

// 初始化TaobaoQimenItemmappingCreateAPIRequest对象
func NewTaobaoQimenItemmappingCreateRequest() *TaobaoQimenItemmappingCreateAPIRequest{
    return &TaobaoQimenItemmappingCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemmappingCreateAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.itemmapping.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemmappingCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenItemmappingCreateAPIRequest) SetRequest(_request *RequestDo) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenItemmappingCreateAPIRequest) GetRequest() *RequestDo {
    return r._request
}
