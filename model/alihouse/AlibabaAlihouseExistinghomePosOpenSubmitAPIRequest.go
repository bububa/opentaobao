package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest pos进件接口 API请求
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
type AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest struct {
	model.Params
	// 同步进件消息
	_syncTradePosOpenDto *SyncTradePosOpenDto
}

// NewAlibabaAlihouseExistinghomePosOpenSubmitRequest 初始化AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomePosOpenSubmitRequest() *AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.pos.open.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncTradePosOpenDto is SyncTradePosOpenDto Setter
// 同步进件消息
func (r *AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) SetSyncTradePosOpenDto(_syncTradePosOpenDto *SyncTradePosOpenDto) error {
	r._syncTradePosOpenDto = _syncTradePosOpenDto
	r.Set("sync_trade_pos_open_dto", _syncTradePosOpenDto)
	return nil
}

// GetSyncTradePosOpenDto SyncTradePosOpenDto Getter
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetSyncTradePosOpenDto() *SyncTradePosOpenDto {
	return r._syncTradePosOpenDto
}
