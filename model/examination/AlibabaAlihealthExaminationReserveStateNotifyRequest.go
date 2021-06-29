package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检状态主动通知 APIRequest
alibaba.alihealth.examination.reserve.state.notify

到了体检当天后，服务商主动通知体检预约状态
*/
type AlibabaAlihealthExaminationReserveStateNotifyRequest struct {
    model.Params

    // 服务商预约凭证
    uniqReserveCode   string 

    // 健康预约凭证
    reserveNumber   string 

    // 体检状态：未到检(exam_not), 已到检(exam_done)； 上门服务中还需以下两种状态：预约确认中（reserve_confirming），预约拒绝（reserve_rejected）；
    reportStatus   string 

    // 到检凭证，exam_done状态下，该字段必填
    checkNo   string 

}

func NewAlibabaAlihealthExaminationReserveStateNotifyRequest() *AlibabaAlihealthExaminationReserveStateNotifyRequest{
    return &AlibabaAlihealthExaminationReserveStateNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.state.notify"
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveStateNotifyRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationReserveStateNotifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveStateNotifyRequest) SetReportStatus(reportStatus string) error {
    r.reportStatus = reportStatus
    r.Set("report_status", reportStatus)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetReportStatus() string {
    return r.reportStatus
}

func (r *AlibabaAlihealthExaminationReserveStateNotifyRequest) SetCheckNo(checkNo string) error {
    r.checkNo = checkNo
    r.Set("check_no", checkNo)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateNotifyRequest) GetCheckNo() string {
    return r.checkNo
}

