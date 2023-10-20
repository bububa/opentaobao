package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamsfserviceworkerqueryidAPIRequest 查询师傅workerid API请求
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
type AlibabamsfserviceworkerqueryidAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 身份证ID
	_idNumber string
}

// NewAlibabamsfserviceworkerqueryidRequest 初始化AlibabamsfserviceworkerqueryidAPIRequest对象
func NewAlibabamsfserviceworkerqueryidRequest() *AlibabamsfserviceworkerqueryidAPIRequest {
	return &AlibabamsfserviceworkerqueryidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamsfserviceworkerqueryidAPIRequest) GetApiMethodName() string {
	return "alibaba.msfservice.worker.queryid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamsfserviceworkerqueryidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamsfserviceworkerqueryidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabamsfserviceworkerqueryidAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabamsfserviceworkerqueryidAPIRequest) GetPhone() string {
	return r._phone
}

// SetIdNumber is IdNumber Setter
// 身份证ID
func (r *AlibabamsfserviceworkerqueryidAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabamsfserviceworkerqueryidAPIRequest) GetIdNumber() string {
	return r._idNumber
}
