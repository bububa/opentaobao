package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuMappingDeleteAPIRequest 货品关系解绑 API请求
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
type AlibabaAscpCnskuMappingDeleteAPIRequest struct {
	model.Params
	// 待解绑的货品关系
	_cnskuRelationDto *CnskuRelationDto
	// 操作信息(不要传null)
	_cnskuRelationOperateOption *CnskuRelationOperateOption
}

// NewAlibabaAscpCnskuMappingDeleteRequest 初始化AlibabaAscpCnskuMappingDeleteAPIRequest对象
func NewAlibabaAscpCnskuMappingDeleteRequest() *AlibabaAscpCnskuMappingDeleteAPIRequest {
	return &AlibabaAscpCnskuMappingDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpCnskuMappingDeleteAPIRequest) Reset() {
	r._cnskuRelationDto = nil
	r._cnskuRelationOperateOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuMappingDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.mapping.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuMappingDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpCnskuMappingDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnskuRelationDto is CnskuRelationDto Setter
// 待解绑的货品关系
func (r *AlibabaAscpCnskuMappingDeleteAPIRequest) SetCnskuRelationDto(_cnskuRelationDto *CnskuRelationDto) error {
	r._cnskuRelationDto = _cnskuRelationDto
	r.Set("cnsku_relation_dto", _cnskuRelationDto)
	return nil
}

// GetCnskuRelationDto CnskuRelationDto Getter
func (r AlibabaAscpCnskuMappingDeleteAPIRequest) GetCnskuRelationDto() *CnskuRelationDto {
	return r._cnskuRelationDto
}

// SetCnskuRelationOperateOption is CnskuRelationOperateOption Setter
// 操作信息(不要传null)
func (r *AlibabaAscpCnskuMappingDeleteAPIRequest) SetCnskuRelationOperateOption(_cnskuRelationOperateOption *CnskuRelationOperateOption) error {
	r._cnskuRelationOperateOption = _cnskuRelationOperateOption
	r.Set("cnsku_relation_operate_option", _cnskuRelationOperateOption)
	return nil
}

// GetCnskuRelationOperateOption CnskuRelationOperateOption Getter
func (r AlibabaAscpCnskuMappingDeleteAPIRequest) GetCnskuRelationOperateOption() *CnskuRelationOperateOption {
	return r._cnskuRelationOperateOption
}

var poolAlibabaAscpCnskuMappingDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpCnskuMappingDeleteRequest()
	},
}

// GetAlibabaAscpCnskuMappingDeleteRequest 从 sync.Pool 获取 AlibabaAscpCnskuMappingDeleteAPIRequest
func GetAlibabaAscpCnskuMappingDeleteAPIRequest() *AlibabaAscpCnskuMappingDeleteAPIRequest {
	return poolAlibabaAscpCnskuMappingDeleteAPIRequest.Get().(*AlibabaAscpCnskuMappingDeleteAPIRequest)
}

// ReleaseAlibabaAscpCnskuMappingDeleteAPIRequest 将 AlibabaAscpCnskuMappingDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpCnskuMappingDeleteAPIRequest(v *AlibabaAscpCnskuMappingDeleteAPIRequest) {
	v.Reset()
	poolAlibabaAscpCnskuMappingDeleteAPIRequest.Put(v)
}
