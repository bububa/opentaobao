package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest 门店状态变更 API请求
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
type AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest struct {
	model.Params
	// dto
	_storeStatusDto *StoreStatusDto
}

// NewAlibabaAlihouseExistinghomeStoreStatusChangeSyncRequest 初始化AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeStoreStatusChangeSyncRequest() *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) Reset() {
	r._storeStatusDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.status.change.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreStatusDto is StoreStatusDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) SetStoreStatusDto(_storeStatusDto *StoreStatusDto) error {
	r._storeStatusDto = _storeStatusDto
	r.Set("store_status_dto", _storeStatusDto)
	return nil
}

// GetStoreStatusDto StoreStatusDto Getter
func (r AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) GetStoreStatusDto() *StoreStatusDto {
	return r._storeStatusDto
}

var poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeStoreStatusChangeSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeStoreStatusChangeSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest
func GetAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest() *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest 将 AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest(v *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest.Put(v)
}
