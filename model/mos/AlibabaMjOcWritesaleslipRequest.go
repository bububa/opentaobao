package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票占库 API请求
alibaba.mj.oc.writesaleslip

开票占库
*/
type AlibabaMjOcWritesaleslipAPIRequest struct {
    model.Params
    // 开票占库入参
    _posSaleOrder   *PosSaleOrderDTO
}

// 初始化AlibabaMjOcWritesaleslipAPIRequest对象
func NewAlibabaMjOcWritesaleslipRequest() *AlibabaMjOcWritesaleslipAPIRequest{
    return &AlibabaMjOcWritesaleslipAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcWritesaleslipAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.writesaleslip"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcWritesaleslipAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosSaleOrder Setter
// 开票占库入参
func (r *AlibabaMjOcWritesaleslipAPIRequest) SetPosSaleOrder(_posSaleOrder *PosSaleOrderDTO) error {
    r._posSaleOrder = _posSaleOrder
    r.Set("pos_sale_order", _posSaleOrder)
    return nil
}

// PosSaleOrder Getter
func (r AlibabaMjOcWritesaleslipAPIRequest) GetPosSaleOrder() *PosSaleOrderDTO {
    return r._posSaleOrder
}
