package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存同步接口 APIRequest
alibaba.ascp.inventory.sync

新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统
*/
type AlibabaAscpInventorySyncRequest struct {
    model.Params

    // 库存同步DTO
    invSingleItemSyncDto   *InvSingleItemSyncDto 

}

func NewAlibabaAscpInventorySyncRequest() *AlibabaAscpInventorySyncRequest{
    return &AlibabaAscpInventorySyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpInventorySyncRequest) GetApiMethodName() string {
    return "alibaba.ascp.inventory.sync"
}

func (r AlibabaAscpInventorySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpInventorySyncRequest) SetInvSingleItemSyncDto(invSingleItemSyncDto *InvSingleItemSyncDto) error {
    r.invSingleItemSyncDto = invSingleItemSyncDto
    r.Set("inv_single_item_sync_dto", invSingleItemSyncDto)
    return nil
}

func (r AlibabaAscpInventorySyncRequest) GetInvSingleItemSyncDto() *InvSingleItemSyncDto {
    return r.invSingleItemSyncDto
}

