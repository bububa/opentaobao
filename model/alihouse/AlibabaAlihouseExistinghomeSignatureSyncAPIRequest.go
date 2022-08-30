package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.signature.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeSignatureSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
