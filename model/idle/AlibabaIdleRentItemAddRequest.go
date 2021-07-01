package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
租赁商品发布 API请求
alibaba.idle.rent.item.add

发布闲鱼租赁商品
*/
type AlibabaIdleRentItemAddAPIRequest struct {
    model.Params
    // 商品信息
    _paramRentItemDTO   *RentItemDTO
}

// 初始化AlibabaIdleRentItemAddAPIRequest对象
func NewAlibabaIdleRentItemAddRequest() *AlibabaIdleRentItemAddAPIRequest{
    return &AlibabaIdleRentItemAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemAddAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemAddAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDTO) error {
    r._paramRentItemDTO = _paramRentItemDTO
    r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
    return nil
}

// ParamRentItemDTO Getter
func (r AlibabaIdleRentItemAddAPIRequest) GetParamRentItemDTO() *RentItemDTO {
    return r._paramRentItemDTO
}
