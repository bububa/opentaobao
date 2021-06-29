package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
租赁商品编辑 APIRequest
alibaba.idle.rent.item.edit

发布闲鱼租赁商品
*/
type AlibabaIdleRentItemEditRequest struct {
    model.Params

    // 商品信息
    paramRentItemDTO   *RentItemDto 

}

func NewAlibabaIdleRentItemEditRequest() *AlibabaIdleRentItemEditRequest{
    return &AlibabaIdleRentItemEditRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentItemEditRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.edit"
}

func (r AlibabaIdleRentItemEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentItemEditRequest) SetParamRentItemDTO(paramRentItemDTO *RentItemDto) error {
    r.paramRentItemDTO = paramRentItemDTO
    r.Set("param_rent_item_d_t_o", paramRentItemDTO)
    return nil
}

func (r AlibabaIdleRentItemEditRequest) GetParamRentItemDTO() *RentItemDto {
    return r.paramRentItemDTO
}

