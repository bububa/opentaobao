package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuAddAPIRequest 货品创建 API请求
// alibaba.ascp.cnsku.add
//
// 供应链中台货品创建接口
type AlibabaAscpCnskuAddAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnskuDto *CnskuDto
	// 选项
	_option *AddCnskuOption
}

// NewAlibabaAscpCnskuAddRequest 初始化AlibabaAscpCnskuAddAPIRequest对象
func NewAlibabaAscpCnskuAddRequest() *AlibabaAscpCnskuAddAPIRequest {
	return &AlibabaAscpCnskuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCnskuDto is CnskuDto Setter
// 待新增的货品
func (r *AlibabaAscpCnskuAddAPIRequest) SetCnskuDto(_cnskuDto *CnskuDto) error {
	r._cnskuDto = _cnskuDto
	r.Set("cnsku_dto", _cnskuDto)
	return nil
}

// GetCnskuDto CnskuDto Getter
func (r AlibabaAscpCnskuAddAPIRequest) GetCnskuDto() *CnskuDto {
	return r._cnskuDto
}

// SetOption is Option Setter
// 选项
func (r *AlibabaAscpCnskuAddAPIRequest) SetOption(_option *AddCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r AlibabaAscpCnskuAddAPIRequest) GetOption() *AddCnskuOption {
	return r._option
}
