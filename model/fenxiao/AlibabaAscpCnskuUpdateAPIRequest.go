package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskuupdateAPIRequest 供应链中台货品修改接口 API请求
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
type AlibabaascpcnskuupdateAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 修改选项
	_option *UpdateCnskuOption
}

// NewAlibabaascpcnskuupdateRequest 初始化AlibabaascpcnskuupdateAPIRequest对象
func NewAlibabaascpcnskuupdateRequest() *AlibabaascpcnskuupdateAPIRequest {
	return &AlibabaascpcnskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpcnskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpcnskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpcnskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnsku is Cnsku Setter
// 待新增的货品
func (r *AlibabaascpcnskuupdateAPIRequest) SetCnsku(_cnsku *CnskuDto) error {
	r._cnsku = _cnsku
	r.Set("cnsku", _cnsku)
	return nil
}

// GetCnsku Cnsku Getter
func (r AlibabaascpcnskuupdateAPIRequest) GetCnsku() *CnskuDto {
	return r._cnsku
}

// SetOption is Option Setter
// 修改选项
func (r *AlibabaascpcnskuupdateAPIRequest) SetOption(_option *UpdateCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r AlibabaascpcnskuupdateAPIRequest) GetOption() *UpdateCnskuOption {
	return r._option
}
