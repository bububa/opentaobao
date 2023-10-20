package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskumappingdeleteAPIRequest 货品关系解绑 API请求
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
type AlibabaascpcnskumappingdeleteAPIRequest struct {
	model.Params
	// 待解绑的货品关系
	_cnskuRelationDto *CnskuRelationDto
	// 操作信息(不要传null)
	_cnskuRelationOperateOption *CnskuRelationOperateOption
}

// NewAlibabaascpcnskumappingdeleteRequest 初始化AlibabaascpcnskumappingdeleteAPIRequest对象
func NewAlibabaascpcnskumappingdeleteRequest() *AlibabaascpcnskumappingdeleteAPIRequest {
	return &AlibabaascpcnskumappingdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpcnskumappingdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.mapping.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpcnskumappingdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpcnskumappingdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnskuRelationDto is CnskuRelationDto Setter
// 待解绑的货品关系
func (r *AlibabaascpcnskumappingdeleteAPIRequest) SetCnskuRelationDto(_cnskuRelationDto *CnskuRelationDto) error {
	r._cnskuRelationDto = _cnskuRelationDto
	r.Set("cnsku_relation_dto", _cnskuRelationDto)
	return nil
}

// GetCnskuRelationDto CnskuRelationDto Getter
func (r AlibabaascpcnskumappingdeleteAPIRequest) GetCnskuRelationDto() *CnskuRelationDto {
	return r._cnskuRelationDto
}

// SetCnskuRelationOperateOption is CnskuRelationOperateOption Setter
// 操作信息(不要传null)
func (r *AlibabaascpcnskumappingdeleteAPIRequest) SetCnskuRelationOperateOption(_cnskuRelationOperateOption *CnskuRelationOperateOption) error {
	r._cnskuRelationOperateOption = _cnskuRelationOperateOption
	r.Set("cnsku_relation_operate_option", _cnskuRelationOperateOption)
	return nil
}

// GetCnskuRelationOperateOption CnskuRelationOperateOption Getter
func (r AlibabaascpcnskumappingdeleteAPIRequest) GetCnskuRelationOperateOption() *CnskuRelationOperateOption {
	return r._cnskuRelationOperateOption
}
