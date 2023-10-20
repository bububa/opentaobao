package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspmailsignuploadAPIRequest IoT售后服务商签收客户邮寄设备附件上传 API请求
// cainiao.iot.ticket.sp.mail.sign.upload
//
// IoT售后服务商签收客户邮寄设备附件上传
type CainiaoiotticketspmailsignuploadAPIRequest struct {
	model.Params
	// 请求参数
	_param *UploadSignVoucherRequest
}

// NewCainiaoiotticketspmailsignuploadRequest 初始化CainiaoiotticketspmailsignuploadAPIRequest对象
func NewCainiaoiotticketspmailsignuploadRequest() *CainiaoiotticketspmailsignuploadAPIRequest {
	return &CainiaoiotticketspmailsignuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketspmailsignuploadAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.mail.sign.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketspmailsignuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketspmailsignuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoiotticketspmailsignuploadAPIRequest) SetParam(_param *UploadSignVoucherRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoiotticketspmailsignuploadAPIRequest) GetParam() *UploadSignVoucherRequest {
	return r._param
}
