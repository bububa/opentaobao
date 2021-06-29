package alihealthcert

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
健康证服务商预约结果通知阿里健康 APIRequest
alibaba.alihealth.examination.reserve.certificate.notify

当ISV执行完健康证预约成功之后， 调用通知阿里健康
*/
type AlibabaAlihealthExaminationReserveCertificateNotifyRequest struct {
    model.Params

    // 服务商预约凭证
    uniqReserveCode   string 

    // 健康预约凭证
    reserveNumber   string 

    // 预约成功（ISSUED），预约失败(FAIL)
    reportStatus   string 

    // 到检凭证，exam_done状态下，该字段必填
    checkNo   string 

    // 预约状态原因，描述预约状态如FAIL时的原因
    statusReason   string 

}

func NewAlibabaAlihealthExaminationReserveCertificateNotifyRequest() *AlibabaAlihealthExaminationReserveCertificateNotifyRequest{
    return &AlibabaAlihealthExaminationReserveCertificateNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.certificate.notify"
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetReportStatus(reportStatus string) error {
    r.reportStatus = reportStatus
    r.Set("report_status", reportStatus)
    return nil
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetReportStatus() string {
    return r.reportStatus
}

func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetCheckNo(checkNo string) error {
    r.checkNo = checkNo
    r.Set("check_no", checkNo)
    return nil
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetCheckNo() string {
    return r.checkNo
}

func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetStatusReason(statusReason string) error {
    r.statusReason = statusReason
    r.Set("status_reason", statusReason)
    return nil
}

func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetStatusReason() string {
    return r.statusReason
}

