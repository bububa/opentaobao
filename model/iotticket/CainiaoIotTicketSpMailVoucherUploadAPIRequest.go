package iotticket

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
