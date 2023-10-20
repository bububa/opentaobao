package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomerenttradebinditemAPIRequest 交易绑定商品 API请求
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
type AlibabaalihouseexistinghomerenttradebinditemAPIRequest struct {
	model.Params
	// 交易商品实体类
	_syncHouseTradeItemDto *SyncHouseTradeItemDto
}

// NewAlibabaalihouseexistinghomerenttradebinditemRequest 初始化AlibabaalihouseexistinghomerenttradebinditemAPIRequest对象
func NewAlibabaalihouseexistinghomerenttradebinditemRequest() *AlibabaalihouseexistinghomerenttradebinditemAPIRequest {
	return &AlibabaalihouseexistinghomerenttradebinditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomerenttradebinditemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.rent.trade.bind.item"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomerenttradebinditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomerenttradebinditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncHouseTradeItemDto is SyncHouseTradeItemDto Setter
// 交易商品实体类
func (r *AlibabaalihouseexistinghomerenttradebinditemAPIRequest) SetSyncHouseTradeItemDto(_syncHouseTradeItemDto *SyncHouseTradeItemDto) error {
	r._syncHouseTradeItemDto = _syncHouseTradeItemDto
	r.Set("sync_house_trade_item_dto", _syncHouseTradeItemDto)
	return nil
}

// GetSyncHouseTradeItemDto SyncHouseTradeItemDto Getter
func (r AlibabaalihouseexistinghomerenttradebinditemAPIRequest) GetSyncHouseTradeItemDto() *SyncHouseTradeItemDto {
	return r._syncHouseTradeItemDto
}
