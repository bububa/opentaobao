package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest 交易绑定商品 API请求
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
type AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest struct {
	model.Params
	// 交易商品实体类
	_syncHouseTradeItemDto *SyncHouseTradeItemDto
}

// NewAlibabaAlihouseExistinghomeRentTradeBindItemRequest 初始化AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest对象
func NewAlibabaAlihouseExistinghomeRentTradeBindItemRequest() *AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest {
	return &AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.rent.trade.bind.item"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncHouseTradeItemDto is SyncHouseTradeItemDto Setter
// 交易商品实体类
func (r *AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest) SetSyncHouseTradeItemDto(_syncHouseTradeItemDto *SyncHouseTradeItemDto) error {
	r._syncHouseTradeItemDto = _syncHouseTradeItemDto
	r.Set("sync_house_trade_item_dto", _syncHouseTradeItemDto)
	return nil
}

// GetSyncHouseTradeItemDto SyncHouseTradeItemDto Getter
func (r AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest) GetSyncHouseTradeItemDto() *SyncHouseTradeItemDto {
	return r._syncHouseTradeItemDto
}
