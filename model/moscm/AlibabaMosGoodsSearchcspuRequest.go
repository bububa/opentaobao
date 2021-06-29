package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
cspu查询 API请求
alibaba.mos.goods.searchcspu

商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
*/
type AlibabaMosGoodsSearchcspuRequest struct {
    model.Params
    // 组合查询对象
    paramCspuCriteria   *CspuCriteria
    // 分页参数
    paramPaginator   *Paginator
}

// 初始化AlibabaMosGoodsSearchcspuRequest对象
func NewAlibabaMosGoodsSearchcspuRequest() *AlibabaMosGoodsSearchcspuRequest{
    return &AlibabaMosGoodsSearchcspuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSearchcspuRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.searchcspu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSearchcspuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCspuCriteria Setter
// 组合查询对象
func (r *AlibabaMosGoodsSearchcspuRequest) SetParamCspuCriteria(paramCspuCriteria *CspuCriteria) error {
    r.paramCspuCriteria = paramCspuCriteria
    r.Set("param_cspu_criteria", paramCspuCriteria)
    return nil
}

// ParamCspuCriteria Getter
func (r AlibabaMosGoodsSearchcspuRequest) GetParamCspuCriteria() *CspuCriteria {
    return r.paramCspuCriteria
}
// ParamPaginator Setter
// 分页参数
func (r *AlibabaMosGoodsSearchcspuRequest) SetParamPaginator(paramPaginator *Paginator) error {
    r.paramPaginator = paramPaginator
    r.Set("param_paginator", paramPaginator)
    return nil
}

// ParamPaginator Getter
func (r AlibabaMosGoodsSearchcspuRequest) GetParamPaginator() *Paginator {
    return r.paramPaginator
}
