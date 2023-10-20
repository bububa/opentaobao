package iotticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMailVoucherUploadAPIRequest 服务商寄出维修件上传凭证信息 API请求
// cainiao.iot.ticket.sp.mail.voucher.upload
//
// IoT售后服务商寄出维修件上传凭证信息
type CainiaoIotTicketSpMailVoucherUploadAPIRequest struct {
	model.Params
	// 请求参数
	_param *CommentTicketTopRequest
}

// NewCainiaoIotTicketSpMailVoucherUploadRequest 初始化CainiaoIotTicketSpMailVoucherUploadAPIRequest对象
func NewCainiaoIotTicketSpMailVoucherUploadRequest() *CainiaoIotTicketSpMailVoucherUploadAPIRequest {
	return &CainiaoIotTicketSpMailVoucherUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketSpMailVoucherUploadAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMailVoucherUploadAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.mail.voucher.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMailVoucherUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketSpMailVoucherUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMailVoucherUploadAPIRequest) SetParam(_param *CommentTicketTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoIotTicketSpMailVoucherUploadAPIRequest) GetParam() *CommentTicketTopRequest {
	return r._param
}

var poolCainiaoIotTicketSpMailVoucherUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketSpMailVoucherUploadRequest()
	},
}

// GetCainiaoIotTicketSpMailVoucherUploadRequest 从 sync.Pool 获取 CainiaoIotTicketSpMailVoucherUploadAPIRequest
func GetCainiaoIotTicketSpMailVoucherUploadAPIRequest() *CainiaoIotTicketSpMailVoucherUploadAPIRequest {
	return poolCainiaoIotTicketSpMailVoucherUploadAPIRequest.Get().(*CainiaoIotTicketSpMailVoucherUploadAPIRequest)
}

// ReleaseCainiaoIotTicketSpMailVoucherUploadAPIRequest 将 CainiaoIotTicketSpMailVoucherUploadAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketSpMailVoucherUploadAPIRequest(v *CainiaoIotTicketSpMailVoucherUploadAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketSpMailVoucherUploadAPIRequest.Put(v)
}
