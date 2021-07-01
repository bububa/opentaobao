package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商签收客户邮寄设备附件上传 API请求
cainiao.iot.ticket.sp.mail.sign.upload

IoT售后服务商签收客户邮寄设备附件上传
*/
type CainiaoIotTicketSpMailSignUploadAPIRequest struct {
    model.Params
    // 请求参数
    _param   *UploadSignVoucherRequest
}

// 初始化CainiaoIotTicketSpMailSignUploadAPIRequest对象
func NewCainiaoIotTicketSpMailSignUploadRequest() *CainiaoIotTicketSpMailSignUploadAPIRequest{
    return &CainiaoIotTicketSpMailSignUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.mail.sign.upload"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMailSignUploadAPIRequest) SetParam(_param *UploadSignVoucherRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpMailSignUploadAPIRequest) GetParam() *UploadSignVoucherRequest {
    return r._param
}
