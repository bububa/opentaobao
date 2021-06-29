package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存同步接口 API请求
alibaba.ascp.inventory.sync

新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统
*/
type AlibabaAscpInventorySyncRequest struct {
    model.Params
    // 库存同步DTO
    _invSingleItemSyncDto   *InvSingleItemSyncDto
}

// 初始化AlibabaAscpInventorySyncRequest对象
func NewAlibabaAscpInventorySyncRequest() *AlibabaAscpInventorySyncRequest{
    return &AlibabaAscpInventorySyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpInventorySyncRequest) GetApiMethodName() string {
    return "alibaba.ascp.inventory.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpInventorySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvSingleItemSyncDto Setter
// 库存同步DTO
func (r *AlibabaAscpInventorySyncRequest) SetInvSingleItemSyncDto(_invSingleItemSyncDto *InvSingleItemSyncDto) error {
    r._invSingleItemSyncDto = _invSingleItemSyncDto
    r.Set("inv_single_item_sync_dto", _invSingleItemSyncDto)
    return nil
}

// InvSingleItemSyncDto Getter
func (r AlibabaAscpInventorySyncRequest) GetInvSingleItemSyncDto() *InvSingleItemSyncDto {
    return r._invSingleItemSyncDto
}
