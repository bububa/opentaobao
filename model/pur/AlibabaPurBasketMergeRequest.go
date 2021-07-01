package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合并购物车 API请求
alibaba.pur.basket.merge

采购商城接入第三方商家合并购物车接口服务
*/
type AlibabaPurBasketMergeAPIRequest struct {
    model.Params
    // 合并购物车入参
    _paramMallMergeCartRequestDTO   *MallMergeCartRequestDTO
}

// 初始化AlibabaPurBasketMergeAPIRequest对象
func NewAlibabaPurBasketMergeRequest() *AlibabaPurBasketMergeAPIRequest{
    return &AlibabaPurBasketMergeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurBasketMergeAPIRequest) GetApiMethodName() string {
    return "alibaba.pur.basket.merge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurBasketMergeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamMallMergeCartRequestDTO Setter
// 合并购物车入参
func (r *AlibabaPurBasketMergeAPIRequest) SetParamMallMergeCartRequestDTO(_paramMallMergeCartRequestDTO *MallMergeCartRequestDTO) error {
    r._paramMallMergeCartRequestDTO = _paramMallMergeCartRequestDTO
    r.Set("param_mall_merge_cart_request_d_t_o", _paramMallMergeCartRequestDTO)
    return nil
}

// ParamMallMergeCartRequestDTO Getter
func (r AlibabaPurBasketMergeAPIRequest) GetParamMallMergeCartRequestDTO() *MallMergeCartRequestDTO {
    return r._paramMallMergeCartRequestDTO
}
