package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomesignaturesyncAPIRequest 二手房电子签章数据同步 API请求
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
type AlibabaalihouseexistinghomesignaturesyncAPIRequest struct {
	model.Params
	// 电子签章实体对象
	_electricSignatureDto *ElectricSignatureDto
}

// NewAlibabaalihouseexistinghomesignaturesyncRequest 初始化AlibabaalihouseexistinghomesignaturesyncAPIRequest对象
func NewAlibabaalihouseexistinghomesignaturesyncRequest() *AlibabaalihouseexistinghomesignaturesyncAPIRequest {
	return &AlibabaalihouseexistinghomesignaturesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomesignaturesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.signature.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomesignaturesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomesignaturesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElectricSignatureDto is ElectricSignatureDto Setter
// 电子签章实体对象
func (r *AlibabaalihouseexistinghomesignaturesyncAPIRequest) SetElectricSignatureDto(_electricSignatureDto *ElectricSignatureDto) error {
	r._electricSignatureDto = _electricSignatureDto
	r.Set("electric_signature_dto", _electricSignatureDto)
	return nil
}

// GetElectricSignatureDto ElectricSignatureDto Getter
func (r AlibabaalihouseexistinghomesignaturesyncAPIRequest) GetElectricSignatureDto() *ElectricSignatureDto {
	return r._electricSignatureDto
}
