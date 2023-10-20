package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest 门店状态变更 API请求
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
type AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest struct {
	model.Params
	// dto
	_storeStatusDto *StoreStatusDto
}

// NewAlibabaalihouseexistinghomestorestatuschangesyncRequest 初始化AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest对象
func NewAlibabaalihouseexistinghomestorestatuschangesyncRequest() *AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest {
	return &AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.status.change.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreStatusDto is StoreStatusDto Setter
// dto
func (r *AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest) SetStoreStatusDto(_storeStatusDto *StoreStatusDto) error {
	r._storeStatusDto = _storeStatusDto
	r.Set("store_status_dto", _storeStatusDto)
	return nil
}

// GetStoreStatusDto StoreStatusDto Getter
func (r AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest) GetStoreStatusDto() *StoreStatusDto {
	return r._storeStatusDto
}
