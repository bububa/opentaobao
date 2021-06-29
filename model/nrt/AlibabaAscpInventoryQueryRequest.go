package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存查询接口 APIRequest
alibaba.ascp.inventory.query

新零售联盟商家库存查询接口，用于商家商品库存数量感知查询
*/
type AlibabaAscpInventoryQueryRequest struct {
    model.Params

    // 系统自动生成
    invSingleItemSyncDto   *InvSingleItemSyncDto 

}

func NewAlibabaAscpInventoryQueryRequest() *AlibabaAscpInventoryQueryRequest{
    return &AlibabaAscpInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpInventoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.inventory.query"
}

func (r AlibabaAscpInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpInventoryQueryRequest) SetInvSingleItemSyncDto(invSingleItemSyncDto *InvSingleItemSyncDto) error {
    r.invSingleItemSyncDto = invSingleItemSyncDto
    r.Set("inv_single_item_sync_dto", invSingleItemSyncDto)
    return nil
}

func (r AlibabaAscpInventoryQueryRequest) GetInvSingleItemSyncDto() *InvSingleItemSyncDto {
    return r.invSingleItemSyncDto
}

