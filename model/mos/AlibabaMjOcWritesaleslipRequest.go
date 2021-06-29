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
type AlibabaMjOcWritesaleslipRequest struct {
    model.Params
    // 开票占库入参
    _posSaleOrder   *PosSaleOrderDto
}

// 初始化AlibabaMjOcWritesaleslipRequest对象
func NewAlibabaMjOcWritesaleslipRequest() *AlibabaMjOcWritesaleslipRequest{
    return &AlibabaMjOcWritesaleslipRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcWritesaleslipRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.writesaleslip"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcWritesaleslipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosSaleOrder Setter
// 开票占库入参
func (r *AlibabaMjOcWritesaleslipRequest) SetPosSaleOrder(_posSaleOrder *PosSaleOrderDto) error {
    r._posSaleOrder = _posSaleOrder
    r.Set("pos_sale_order", _posSaleOrder)
    return nil
}

// PosSaleOrder Getter
func (r AlibabaMjOcWritesaleslipRequest) GetPosSaleOrder() *PosSaleOrderDto {
    return r._posSaleOrder
}
