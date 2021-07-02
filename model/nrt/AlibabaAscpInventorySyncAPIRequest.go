package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpInventorySyncAPIRequest 库存同步接口 API请求
// alibaba.ascp.inventory.sync
//
// 新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统
type AlibabaAscpInventorySyncAPIRequest struct {
	model.Params
	// 库存同步DTO
	_invSingleItemSyncDto *InvSingleItemSyncDto
}

// NewAlibabaAscpInventorySyncRequest 初始化AlibabaAscpInventorySyncAPIRequest对象
func NewAlibabaAscpInventorySyncRequest() *AlibabaAscpInventorySyncAPIRequest {
	return &AlibabaAscpInventorySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpInventorySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.inventory.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpInventorySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvSingleItemSyncDto Setter
// 库存同步DTO
func (r *AlibabaAscpInventorySyncAPIRequest) SetInvSingleItemSyncDto(_invSingleItemSyncDto *InvSingleItemSyncDto) error {
	r._invSingleItemSyncDto = _invSingleItemSyncDto
	r.Set("inv_single_item_sync_dto", _invSingleItemSyncDto)
	return nil
}

// Get InvSingleItemSyncDto Getter
func (r AlibabaAscpInventorySyncAPIRequest) GetInvSingleItemSyncDto() *InvSingleItemSyncDto {
	return r._invSingleItemSyncDto
}
