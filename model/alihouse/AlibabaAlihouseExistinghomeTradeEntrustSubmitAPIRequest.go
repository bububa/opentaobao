package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest 交易委托信息更新接口 API请求
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
type AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest struct {
	model.Params
	// 更新交易委托实体类
	_syncUpdateTradeEntrustDto *SyncUpdateTradeEntrustDto
}

// NewAlibabaAlihouseExistinghomeTradeEntrustSubmitRequest 初始化AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomeTradeEntrustSubmitRequest() *AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.trade.entrust.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncUpdateTradeEntrustDto is SyncUpdateTradeEntrustDto Setter
// 更新交易委托实体类
func (r *AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest) SetSyncUpdateTradeEntrustDto(_syncUpdateTradeEntrustDto *SyncUpdateTradeEntrustDto) error {
	r._syncUpdateTradeEntrustDto = _syncUpdateTradeEntrustDto
	r.Set("sync_update_trade_entrust_dto", _syncUpdateTradeEntrustDto)
	return nil
}

// GetSyncUpdateTradeEntrustDto SyncUpdateTradeEntrustDto Getter
func (r AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest) GetSyncUpdateTradeEntrustDto() *SyncUpdateTradeEntrustDto {
	return r._syncUpdateTradeEntrustDto
}
