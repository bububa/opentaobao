package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuUpdateAPIRequest 供应链中台货品修改接口 API请求
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
type AlibabaAscpCnskuUpdateAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 修改选项
	_option *UpdateCnskuOption
}

// NewAlibabaAscpCnskuUpdateRequest 初始化AlibabaAscpCnskuUpdateAPIRequest对象
func NewAlibabaAscpCnskuUpdateRequest() *AlibabaAscpCnskuUpdateAPIRequest {
	return &AlibabaAscpCnskuUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpCnskuUpdateAPIRequest) Reset() {
	r._cnsku = nil
	r._option = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpCnskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnsku is Cnsku Setter
// 待新增的货品
func (r *AlibabaAscpCnskuUpdateAPIRequest) SetCnsku(_cnsku *CnskuDto) error {
	r._cnsku = _cnsku
	r.Set("cnsku", _cnsku)
	return nil
}

// GetCnsku Cnsku Getter
func (r AlibabaAscpCnskuUpdateAPIRequest) GetCnsku() *CnskuDto {
	return r._cnsku
}

// SetOption is Option Setter
// 修改选项
func (r *AlibabaAscpCnskuUpdateAPIRequest) SetOption(_option *UpdateCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r AlibabaAscpCnskuUpdateAPIRequest) GetOption() *UpdateCnskuOption {
	return r._option
}

var poolAlibabaAscpCnskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpCnskuUpdateRequest()
	},
}

// GetAlibabaAscpCnskuUpdateRequest 从 sync.Pool 获取 AlibabaAscpCnskuUpdateAPIRequest
func GetAlibabaAscpCnskuUpdateAPIRequest() *AlibabaAscpCnskuUpdateAPIRequest {
	return poolAlibabaAscpCnskuUpdateAPIRequest.Get().(*AlibabaAscpCnskuUpdateAPIRequest)
}

// ReleaseAlibabaAscpCnskuUpdateAPIRequest 将 AlibabaAscpCnskuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpCnskuUpdateAPIRequest(v *AlibabaAscpCnskuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAscpCnskuUpdateAPIRequest.Put(v)
}
