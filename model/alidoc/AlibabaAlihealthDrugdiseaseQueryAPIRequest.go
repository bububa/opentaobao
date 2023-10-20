package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdiseasequeryAPIRequest 药品诊断查询 API请求
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
type AlibabaalihealthdrugdiseasequeryAPIRequest struct {
	model.Params
	// 入参
	_spuDiseaseQueryDto *SpuDiseaseQueryDto
}

// NewAlibabaalihealthdrugdiseasequeryRequest 初始化AlibabaalihealthdrugdiseasequeryAPIRequest对象
func NewAlibabaalihealthdrugdiseasequeryRequest() *AlibabaalihealthdrugdiseasequeryAPIRequest {
	return &AlibabaalihealthdrugdiseasequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdiseasequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugdisease.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdiseasequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdiseasequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpuDiseaseQueryDto is SpuDiseaseQueryDto Setter
// 入参
func (r *AlibabaalihealthdrugdiseasequeryAPIRequest) SetSpuDiseaseQueryDto(_spuDiseaseQueryDto *SpuDiseaseQueryDto) error {
	r._spuDiseaseQueryDto = _spuDiseaseQueryDto
	r.Set("spu_disease_query_dto", _spuDiseaseQueryDto)
	return nil
}

// GetSpuDiseaseQueryDto SpuDiseaseQueryDto Getter
func (r AlibabaalihealthdrugdiseasequeryAPIRequest) GetSpuDiseaseQueryDto() *SpuDiseaseQueryDto {
	return r._spuDiseaseQueryDto
}
