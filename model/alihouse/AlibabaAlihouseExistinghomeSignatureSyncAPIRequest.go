package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSignatureSyncAPIRequest 二手房电子签章数据同步 API请求
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
type AlibabaAlihouseExistinghomeSignatureSyncAPIRequest struct {
	model.Params
	// 电子签章实体对象
	_electricSignatureDto *ElectricSignatureDto
}

// NewAlibabaAlihouseExistinghomeSignatureSyncRequest 初始化AlibabaAlihouseExistinghomeSignatureSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeSignatureSyncRequest() *AlibabaAlihouseExistinghomeSignatureSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeSignatureSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) Reset() {
	r._electricSignatureDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.signature.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElectricSignatureDto is ElectricSignatureDto Setter
// 电子签章实体对象
func (r *AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) SetElectricSignatureDto(_electricSignatureDto *ElectricSignatureDto) error {
	r._electricSignatureDto = _electricSignatureDto
	r.Set("electric_signature_dto", _electricSignatureDto)
	return nil
}

// GetElectricSignatureDto ElectricSignatureDto Getter
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetElectricSignatureDto() *ElectricSignatureDto {
	return r._electricSignatureDto
}

var poolAlibabaAlihouseExistinghomeSignatureSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeSignatureSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeSignatureSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSignatureSyncAPIRequest
func GetAlibabaAlihouseExistinghomeSignatureSyncAPIRequest() *AlibabaAlihouseExistinghomeSignatureSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeSignatureSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeSignatureSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeSignatureSyncAPIRequest 将 AlibabaAlihouseExistinghomeSignatureSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSignatureSyncAPIRequest(v *AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSignatureSyncAPIRequest.Put(v)
}
