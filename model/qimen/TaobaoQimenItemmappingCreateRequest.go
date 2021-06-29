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
type TaobaoQimenItemmappingCreateRequest struct {
    model.Params
    // 
    request   *RequestDO
}

// 初始化TaobaoQimenItemmappingCreateRequest对象
func NewTaobaoQimenItemmappingCreateRequest() *TaobaoQimenItemmappingCreateRequest{
    return &TaobaoQimenItemmappingCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemmappingCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.itemmapping.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemmappingCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenItemmappingCreateRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenItemmappingCreateRequest) GetRequest() *RequestDO {
    return r.request
}
