package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeApartmentOuteridAPIRequest 公寓更新outerid API请求
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
type AlibabaAlihouseNewhomeApartmentOuteridAPIRequest struct {
	model.Params
	// 公寓outerid
	_outerId string
	// 公寓ecode
	_eCode string
}

// NewAlibabaAlihouseNewhomeApartmentOuteridRequest 初始化AlibabaAlihouseNewhomeApartmentOuteridAPIRequest对象
func NewAlibabaAlihouseNewhomeApartmentOuteridRequest() *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest {
	return &AlibabaAlihouseNewhomeApartmentOuteridAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) Reset() {
	r._outerId = ""
	r._eCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.apartment.outerid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 公寓outerid
func (r *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetECode is ECode Setter
// 公寓ecode
func (r *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) SetECode(_eCode string) error {
	r._eCode = _eCode
	r.Set("e_code", _eCode)
	return nil
}

// GetECode ECode Getter
func (r AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) GetECode() string {
	return r._eCode
}

var poolAlibabaAlihouseNewhomeApartmentOuteridAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeApartmentOuteridRequest()
	},
}

// GetAlibabaAlihouseNewhomeApartmentOuteridRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeApartmentOuteridAPIRequest
func GetAlibabaAlihouseNewhomeApartmentOuteridAPIRequest() *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest {
	return poolAlibabaAlihouseNewhomeApartmentOuteridAPIRequest.Get().(*AlibabaAlihouseNewhomeApartmentOuteridAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeApartmentOuteridAPIRequest 将 AlibabaAlihouseNewhomeApartmentOuteridAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeApartmentOuteridAPIRequest(v *AlibabaAlihouseNewhomeApartmentOuteridAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeApartmentOuteridAPIRequest.Put(v)
}
