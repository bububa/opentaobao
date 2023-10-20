package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest 门店保证金余额同步 API请求
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
type AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest struct {
	model.Params
	// dto
	_storeBailDto *StoreBailDto
}

// NewAlibabaAlihouseExistinghomeStoreBailSyncRequest 初始化AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeStoreBailSyncRequest() *AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) Reset() {
	r._storeBailDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.bail.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreBailDto is StoreBailDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) SetStoreBailDto(_storeBailDto *StoreBailDto) error {
	r._storeBailDto = _storeBailDto
	r.Set("store_bail_dto", _storeBailDto)
	return nil
}

// GetStoreBailDto StoreBailDto Getter
func (r AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) GetStoreBailDto() *StoreBailDto {
	return r._storeBailDto
}

var poolAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeStoreBailSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeStoreBailSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest
func GetAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest() *AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest 将 AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest(v *AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreBailSyncAPIRequest.Put(v)
}
