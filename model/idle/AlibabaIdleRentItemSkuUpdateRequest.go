package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新/增加sku信息 API请求
alibaba.idle.rent.item.sku.update

更新/增加sku信息
*/
type AlibabaIdleRentItemSkuUpdateRequest struct {
    model.Params
    // sku信息，更新后skuId保持不变
    _sku   *ItemSkuDto
}

// 初始化AlibabaIdleRentItemSkuUpdateRequest对象
func NewAlibabaIdleRentItemSkuUpdateRequest() *AlibabaIdleRentItemSkuUpdateRequest{
    return &AlibabaIdleRentItemSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemSkuUpdateRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Sku Setter
// sku信息，更新后skuId保持不变
func (r *AlibabaIdleRentItemSkuUpdateRequest) SetSku(_sku *ItemSkuDto) error {
    r._sku = _sku
    r.Set("sku", _sku)
    return nil
}

// Sku Getter
func (r AlibabaIdleRentItemSkuUpdateRequest) GetSku() *ItemSkuDto {
    return r._sku
}
