package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存查询接口 API请求
alibaba.ascp.inventory.query

新零售联盟商家库存查询接口，用于商家商品库存数量感知查询
*/
type AlibabaAscpInventoryQueryRequest struct {
    model.Params
    // 系统自动生成
    _invSingleItemSyncDto   *InvSingleItemSyncDto
}

// 初始化AlibabaAscpInventoryQueryRequest对象
func NewAlibabaAscpInventoryQueryRequest() *AlibabaAscpInventoryQueryRequest{
    return &AlibabaAscpInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpInventoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvSingleItemSyncDto Setter
// 系统自动生成
func (r *AlibabaAscpInventoryQueryRequest) SetInvSingleItemSyncDto(_invSingleItemSyncDto *InvSingleItemSyncDto) error {
    r._invSingleItemSyncDto = _invSingleItemSyncDto
    r.Set("inv_single_item_sync_dto", _invSingleItemSyncDto)
    return nil
}

// InvSingleItemSyncDto Getter
func (r AlibabaAscpInventoryQueryRequest) GetInvSingleItemSyncDto() *InvSingleItemSyncDto {
    return r._invSingleItemSyncDto
}
