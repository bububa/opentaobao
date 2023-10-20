package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskumodifyAPIRequest 供应链中台货品修改接口 API请求
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
type AlibabaascpcnskumodifyAPIRequest struct {
	model.Params
	// 待修改的货品
	_cnskuDto *CnskuDto
	// 选项
	_updateCnskuOption *UpdateCnskuOption
}

// NewAlibabaascpcnskumodifyRequest 初始化AlibabaascpcnskumodifyAPIRequest对象
func NewAlibabaascpcnskumodifyRequest() *AlibabaascpcnskumodifyAPIRequest {
	return &AlibabaascpcnskumodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpcnskumodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpcnskumodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpcnskumodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnskuDto is CnskuDto Setter
// 待修改的货品
func (r *AlibabaascpcnskumodifyAPIRequest) SetCnskuDto(_cnskuDto *CnskuDto) error {
	r._cnskuDto = _cnskuDto
	r.Set("cnsku_dto", _cnskuDto)
	return nil
}

// GetCnskuDto CnskuDto Getter
func (r AlibabaascpcnskumodifyAPIRequest) GetCnskuDto() *CnskuDto {
	return r._cnskuDto
}

// SetUpdateCnskuOption is UpdateCnskuOption Setter
// 选项
func (r *AlibabaascpcnskumodifyAPIRequest) SetUpdateCnskuOption(_updateCnskuOption *UpdateCnskuOption) error {
	r._updateCnskuOption = _updateCnskuOption
	r.Set("update_cnsku_option", _updateCnskuOption)
	return nil
}

// GetUpdateCnskuOption UpdateCnskuOption Getter
func (r AlibabaascpcnskumodifyAPIRequest) GetUpdateCnskuOption() *UpdateCnskuOption {
	return r._updateCnskuOption
}
