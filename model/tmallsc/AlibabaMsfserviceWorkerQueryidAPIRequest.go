package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMsfserviceWorkerQueryidAPIRequest 查询师傅workerid API请求
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
type AlibabaMsfserviceWorkerQueryidAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 身份证ID
	_idNumber string
}

// NewAlibabaMsfserviceWorkerQueryidRequest 初始化AlibabaMsfserviceWorkerQueryidAPIRequest对象
func NewAlibabaMsfserviceWorkerQueryidRequest() *AlibabaMsfserviceWorkerQueryidAPIRequest {
	return &AlibabaMsfserviceWorkerQueryidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMsfserviceWorkerQueryidAPIRequest) GetApiMethodName() string {
	return "alibaba.msfservice.worker.queryid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMsfserviceWorkerQueryidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMsfserviceWorkerQueryidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaMsfserviceWorkerQueryidAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaMsfserviceWorkerQueryidAPIRequest) GetPhone() string {
	return r._phone
}

// SetIdNumber is IdNumber Setter
// 身份证ID
func (r *AlibabaMsfserviceWorkerQueryidAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabaMsfserviceWorkerQueryidAPIRequest) GetIdNumber() string {
	return r._idNumber
}
