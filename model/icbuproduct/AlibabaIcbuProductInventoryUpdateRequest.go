package icbuproduct

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu商品库存更新 API请求
alibaba.icbu.product.inventory.update

更新库存信息
*/
type AlibabaIcbuProductInventoryUpdateRequest struct {
    model.Params
    // 更新请求
    _requestParam   *ProductInventoryRequest
}

// 初始化AlibabaIcbuProductInventoryUpdateRequest对象
func NewAlibabaIcbuProductInventoryUpdateRequest() *AlibabaIcbuProductInventoryUpdateRequest{
    return &AlibabaIcbuProductInventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductInventoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.inventory.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductInventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestParam Setter
// 更新请求
func (r *AlibabaIcbuProductInventoryUpdateRequest) SetRequestParam(_requestParam *ProductInventoryRequest) error {
    r._requestParam = _requestParam
    r.Set("request_param", _requestParam)
    return nil
}

// RequestParam Getter
func (r AlibabaIcbuProductInventoryUpdateRequest) GetRequestParam() *ProductInventoryRequest {
    return r._requestParam
}
