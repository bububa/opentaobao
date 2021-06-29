package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
可售库存查询 API请求
alibaba.mos.goods.inventory.getinventorys

查询商品的可售、在库和占库数量
*/
type AlibabaMosGoodsInventoryGetinventorysRequest struct {
    model.Params
    // 查询对象
    _paramVirtualInventoryQueryDto   *VirtualInventoryQueryDto
}

// 初始化AlibabaMosGoodsInventoryGetinventorysRequest对象
func NewAlibabaMosGoodsInventoryGetinventorysRequest() *AlibabaMosGoodsInventoryGetinventorysRequest{
    return &AlibabaMosGoodsInventoryGetinventorysRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.inventory.getinventorys"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamVirtualInventoryQueryDto Setter
// 查询对象
func (r *AlibabaMosGoodsInventoryGetinventorysRequest) SetParamVirtualInventoryQueryDto(_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto) error {
    r._paramVirtualInventoryQueryDto = _paramVirtualInventoryQueryDto
    r.Set("param_virtual_inventory_query_dto", _paramVirtualInventoryQueryDto)
    return nil
}

// ParamVirtualInventoryQueryDto Getter
func (r AlibabaMosGoodsInventoryGetinventorysRequest) GetParamVirtualInventoryQueryDto() *VirtualInventoryQueryDto {
    return r._paramVirtualInventoryQueryDto
}
