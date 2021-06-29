package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
租赁商品编辑 API请求
alibaba.idle.rent.item.edit

发布闲鱼租赁商品
*/
type AlibabaIdleRentItemEditRequest struct {
    model.Params
    // 商品信息
    _paramRentItemDTO   *RentItemDto
}

// 初始化AlibabaIdleRentItemEditRequest对象
func NewAlibabaIdleRentItemEditRequest() *AlibabaIdleRentItemEditRequest{
    return &AlibabaIdleRentItemEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemEditRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.edit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemEditRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
    r._paramRentItemDTO = _paramRentItemDTO
    r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
    return nil
}

// ParamRentItemDTO Getter
func (r AlibabaIdleRentItemEditRequest) GetParamRentItemDTO() *RentItemDto {
    return r._paramRentItemDTO
}
