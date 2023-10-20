package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugdiseaseQueryAPIRequest 药品诊断查询 API请求
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
type AlibabaAlihealthDrugdiseaseQueryAPIRequest struct {
	model.Params
	// 入参
	_spuDiseaseQueryDto *SpuDiseaseQueryDto
}

// NewAlibabaAlihealthDrugdiseaseQueryRequest 初始化AlibabaAlihealthDrugdiseaseQueryAPIRequest对象
func NewAlibabaAlihealthDrugdiseaseQueryRequest() *AlibabaAlihealthDrugdiseaseQueryAPIRequest {
	return &AlibabaAlihealthDrugdiseaseQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugdiseaseQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugdisease.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugdiseaseQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugdiseaseQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpuDiseaseQueryDto is SpuDiseaseQueryDto Setter
// 入参
func (r *AlibabaAlihealthDrugdiseaseQueryAPIRequest) SetSpuDiseaseQueryDto(_spuDiseaseQueryDto *SpuDiseaseQueryDto) error {
	r._spuDiseaseQueryDto = _spuDiseaseQueryDto
	r.Set("spu_disease_query_dto", _spuDiseaseQueryDto)
	return nil
}

// GetSpuDiseaseQueryDto SpuDiseaseQueryDto Getter
func (r AlibabaAlihealthDrugdiseaseQueryAPIRequest) GetSpuDiseaseQueryDto() *SpuDiseaseQueryDto {
	return r._spuDiseaseQueryDto
}
