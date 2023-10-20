package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskuaddAPIRequest 货品创建 API请求
// alibaba.ascp.cnsku.add
//
// 供应链中台货品创建接口
type AlibabaascpcnskuaddAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnskuDto *CnskuDto
	// 选项
	_option *AddCnskuOption
}

// NewAlibabaascpcnskuaddRequest 初始化AlibabaascpcnskuaddAPIRequest对象
func NewAlibabaascpcnskuaddRequest() *AlibabaascpcnskuaddAPIRequest {
	return &AlibabaascpcnskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpcnskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpcnskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpcnskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnskuDto is CnskuDto Setter
// 待新增的货品
func (r *AlibabaascpcnskuaddAPIRequest) SetCnskuDto(_cnskuDto *CnskuDto) error {
	r._cnskuDto = _cnskuDto
	r.Set("cnsku_dto", _cnskuDto)
	return nil
}

// GetCnskuDto CnskuDto Getter
func (r AlibabaascpcnskuaddAPIRequest) GetCnskuDto() *CnskuDto {
	return r._cnskuDto
}

// SetOption is Option Setter
// 选项
func (r *AlibabaascpcnskuaddAPIRequest) SetOption(_option *AddCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r AlibabaascpcnskuaddAPIRequest) GetOption() *AddCnskuOption {
	return r._option
}
