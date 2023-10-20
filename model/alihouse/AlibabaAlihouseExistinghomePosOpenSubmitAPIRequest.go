package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeposopensubmitAPIRequest pos进件接口 API请求
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
type AlibabaalihouseexistinghomeposopensubmitAPIRequest struct {
	model.Params
	// 同步进件消息
	_syncTradePosOpenDto *SyncTradePosOpenDto
}

// NewAlibabaalihouseexistinghomeposopensubmitRequest 初始化AlibabaalihouseexistinghomeposopensubmitAPIRequest对象
func NewAlibabaalihouseexistinghomeposopensubmitRequest() *AlibabaalihouseexistinghomeposopensubmitAPIRequest {
	return &AlibabaalihouseexistinghomeposopensubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeposopensubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.pos.open.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeposopensubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeposopensubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncTradePosOpenDto is SyncTradePosOpenDto Setter
// 同步进件消息
func (r *AlibabaalihouseexistinghomeposopensubmitAPIRequest) SetSyncTradePosOpenDto(_syncTradePosOpenDto *SyncTradePosOpenDto) error {
	r._syncTradePosOpenDto = _syncTradePosOpenDto
	r.Set("sync_trade_pos_open_dto", _syncTradePosOpenDto)
	return nil
}

// GetSyncTradePosOpenDto SyncTradePosOpenDto Getter
func (r AlibabaalihouseexistinghomeposopensubmitAPIRequest) GetSyncTradePosOpenDto() *SyncTradePosOpenDto {
	return r._syncTradePosOpenDto
}
