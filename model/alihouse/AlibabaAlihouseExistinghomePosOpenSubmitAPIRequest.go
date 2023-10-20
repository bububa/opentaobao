package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) Reset() {
	r._syncTradePosOpenDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.pos.open.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomePosOpenSubmitRequest()
	},
}

// GetAlibabaAlihouseExistinghomePosOpenSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest
func GetAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest() *AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest {
	return poolAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest.Get().(*AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest 将 AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest(v *AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomePosOpenSubmitAPIRequest.Put(v)
}
