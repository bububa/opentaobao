package iotticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMailSignUploadAPIRequest IoT售后服务商签收客户邮寄设备附件上传 API请求
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
type CainiaoIotTicketSpMailSignUploadAPIRequest struct {
	model.Params
	// 请求参数
	_param *UploadSignVoucherRequest
}

// NewCainiaoIotTicketSpMailSignUploadRequest 初始化CainiaoIotTicketSpMailSignUploadAPIRequest对象
func NewCainiaoIotTicketSpMailSignUploadRequest() *CainiaoIotTicketSpMailSignUploadAPIRequest {
	return &CainiaoIotTicketSpMailSignUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketSpMailSignUploadAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.mail.sign.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMailSignUploadAPIRequest) SetParam(_param *UploadSignVoucherRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetParam() *UploadSignVoucherRequest {
	return r._param
}

var poolCainiaoIotTicketSpMailSignUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketSpMailSignUploadRequest()
	},
}

// GetCainiaoIotTicketSpMailSignUploadRequest 从 sync.Pool 获取 CainiaoIotTicketSpMailSignUploadAPIRequest
func GetCainiaoIotTicketSpMailSignUploadAPIRequest() *CainiaoIotTicketSpMailSignUploadAPIRequest {
	return poolCainiaoIotTicketSpMailSignUploadAPIRequest.Get().(*CainiaoIotTicketSpMailSignUploadAPIRequest)
}

// ReleaseCainiaoIotTicketSpMailSignUploadAPIRequest 将 CainiaoIotTicketSpMailSignUploadAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketSpMailSignUploadAPIRequest(v *CainiaoIotTicketSpMailSignUploadAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketSpMailSignUploadAPIRequest.Put(v)
}
