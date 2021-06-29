package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商签收客户邮寄设备附件上传 APIRequest
cainiao.iot.ticket.sp.mail.sign.upload

IoT售后服务商签收客户邮寄设备附件上传
*/
type CainiaoIotTicketSpMailSignUploadRequest struct {
    model.Params

    // 请求参数
    param   *UploadSignVoucherRequest 

}

func NewCainiaoIotTicketSpMailSignUploadRequest() *CainiaoIotTicketSpMailSignUploadRequest{
    return &CainiaoIotTicketSpMailSignUploadRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpMailSignUploadRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.mail.sign.upload"
}

func (r CainiaoIotTicketSpMailSignUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpMailSignUploadRequest) SetParam(param *UploadSignVoucherRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpMailSignUploadRequest) GetParam() *UploadSignVoucherRequest {
    return r.param
}

