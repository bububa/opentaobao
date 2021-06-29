package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
可售库存查询 APIRequest
alibaba.mos.goods.inventory.getinventorys

查询商品的可售、在库和占库数量
*/
type AlibabaMosGoodsInventoryGetinventorysRequest struct {
    model.Params

    // 查询对象
    paramVirtualInventoryQueryDto   *VirtualInventoryQueryDto 

}

func NewAlibabaMosGoodsInventoryGetinventorysRequest() *AlibabaMosGoodsInventoryGetinventorysRequest{
    return &AlibabaMosGoodsInventoryGetinventorysRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.inventory.getinventorys"
}

func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosGoodsInventoryGetinventorysRequest) SetParamVirtualInventoryQueryDto(paramVirtualInventoryQueryDto *VirtualInventoryQueryDto) error {
    r.paramVirtualInventoryQueryDto = paramVirtualInventoryQueryDto
    r.Set("param_virtual_inventory_query_dto", paramVirtualInventoryQueryDto)
    return nil
}

func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetParamVirtualInventoryQueryDto() *VirtualInventoryQueryDto {
    return r.paramVirtualInventoryQueryDto
}

