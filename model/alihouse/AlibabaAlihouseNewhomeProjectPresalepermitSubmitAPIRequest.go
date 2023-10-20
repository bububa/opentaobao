package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest 提交预售证 API请求
// alibaba.alihouse.newhome.project.presalepermit.submit
//
// 提交楼盘预售证信息
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest struct {
	model.Params
	// 预售证对象
	_preSalePermitDto *PreSalePermitDto
}

// NewAlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest 初始化AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.presalepermit.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreSalePermitDto is PreSalePermitDto Setter
// 预售证对象
func (r *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) SetPreSalePermitDto(_preSalePermitDto *PreSalePermitDto) error {
	r._preSalePermitDto = _preSalePermitDto
	r.Set("pre_sale_permit_dto", _preSalePermitDto)
	return nil
}

// GetPreSalePermitDto PreSalePermitDto Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetPreSalePermitDto() *PreSalePermitDto {
	return r._preSalePermitDto
}
