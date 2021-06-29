package iotticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iotticket"
)

/* 
IoT售后服务商签收客户邮寄设备附件上传 
cainiao.iot.ticket.sp.mail.sign.upload

IoT售后服务商签收客户邮寄设备附件上传
*/
func CainiaoIotTicketSpMailSignUpload(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMailSignUploadRequest, session string) (*iotticket.CainiaoIotTicketSpMailSignUploadAPIResponse, error) {
    var resp iotticket.CainiaoIotTicketSpMailSignUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
