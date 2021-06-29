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
type AlibabaIdleRentItemAddRequest struct {
    model.Params
    // 商品信息
    paramRentItemDTO   *RentItemDto
}

// 初始化AlibabaIdleRentItemAddRequest对象
func NewAlibabaIdleRentItemAddRequest() *AlibabaIdleRentItemAddRequest{
    return &AlibabaIdleRentItemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemAddRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemAddRequest) SetParamRentItemDTO(paramRentItemDTO *RentItemDto) error {
    r.paramRentItemDTO = paramRentItemDTO
    r.Set("param_rent_item_d_t_o", paramRentItemDTO)
    return nil
}

// ParamRentItemDTO Getter
func (r AlibabaIdleRentItemAddRequest) GetParamRentItemDTO() *RentItemDto {
    return r.paramRentItemDTO
}
