package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合并购物车 APIRequest
alibaba.pur.basket.merge

采购商城接入第三方商家合并购物车接口服务
*/
type AlibabaPurBasketMergeRequest struct {
    model.Params

    // 合并购物车入参
    paramMallMergeCartRequestDTO   *MallMergeCartRequestDto 

}

func NewAlibabaPurBasketMergeRequest() *AlibabaPurBasketMergeRequest{
    return &AlibabaPurBasketMergeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurBasketMergeRequest) GetApiMethodName() string {
    return "alibaba.pur.basket.merge"
}

func (r AlibabaPurBasketMergeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurBasketMergeRequest) SetParamMallMergeCartRequestDTO(paramMallMergeCartRequestDTO *MallMergeCartRequestDto) error {
    r.paramMallMergeCartRequestDTO = paramMallMergeCartRequestDTO
    r.Set("param_mall_merge_cart_request_d_t_o", paramMallMergeCartRequestDTO)
    return nil
}

func (r AlibabaPurBasketMergeRequest) GetParamMallMergeCartRequestDTO() *MallMergeCartRequestDto {
    return r.paramMallMergeCartRequestDTO
}

