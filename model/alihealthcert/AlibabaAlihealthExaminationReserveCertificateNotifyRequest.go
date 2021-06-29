package alihealthcert

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
健康证服务商预约结果通知阿里健康 API请求
alibaba.alihealth.examination.reserve.certificate.notify

当ISV执行完健康证预约成功之后， 调用通知阿里健康
*/
type AlibabaAlihealthExaminationReserveCertificateNotifyRequest struct {
    model.Params
    // 服务商预约凭证
    _uniqReserveCode   string
    // 健康预约凭证
    _reserveNumber   string
    // 预约成功（ISSUED），预约失败(FAIL)
    _reportStatus   string
    // 到检凭证，exam_done状态下，该字段必填
    _checkNo   string
    // 预约状态原因，描述预约状态如FAIL时的原因
    _statusReason   string
}

// 初始化AlibabaAlihealthExaminationReserveCertificateNotifyRequest对象
func NewAlibabaAlihealthExaminationReserveCertificateNotifyRequest() *AlibabaAlihealthExaminationReserveCertificateNotifyRequest{
    return &AlibabaAlihealthExaminationReserveCertificateNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.certificate.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetUniqReserveCode(_uniqReserveCode string) error {
    r._uniqReserveCode = _uniqReserveCode
    r.Set("uniq_reserve_code", _uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetUniqReserveCode() string {
    return r._uniqReserveCode
}
// ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// ReportStatus Setter
// 预约成功（ISSUED），预约失败(FAIL)
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetReportStatus(_reportStatus string) error {
    r._reportStatus = _reportStatus
    r.Set("report_status", _reportStatus)
    return nil
}

// ReportStatus Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetReportStatus() string {
    return r._reportStatus
}
// CheckNo Setter
// 到检凭证，exam_done状态下，该字段必填
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetCheckNo(_checkNo string) error {
    r._checkNo = _checkNo
    r.Set("check_no", _checkNo)
    return nil
}

// CheckNo Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetCheckNo() string {
    return r._checkNo
}
// StatusReason Setter
// 预约状态原因，描述预约状态如FAIL时的原因
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyRequest) SetStatusReason(_statusReason string) error {
    r._statusReason = _statusReason
    r.Set("status_reason", _statusReason)
    return nil
}

// StatusReason Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyRequest) GetStatusReason() string {
    return r._statusReason
}
