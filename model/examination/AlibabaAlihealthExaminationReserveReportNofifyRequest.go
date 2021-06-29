package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商主动通知体检报告 APIRequest
alibaba.alihealth.examination.reserve.report.nofify

服务商主动回传用户的体检报告数据
*/
type AlibabaAlihealthExaminationReserveReportNofifyRequest struct {
    model.Params

    // 服务商预约凭证
    uniqReserveCode   string 

    // 服务商到检编号
    checkNo   string 

    // 健康预约凭证
    reserveNumber   string 

    // 报告通知类型，传1即可
    type   string 

    // pdf文件的二进制base64编码字符串
    content   string 

}

func NewAlibabaAlihealthExaminationReserveReportNofifyRequest() *AlibabaAlihealthExaminationReserveReportNofifyRequest{
    return &AlibabaAlihealthExaminationReserveReportNofifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.report.nofify"
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveReportNofifyRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationReserveReportNofifyRequest) SetCheckNo(checkNo string) error {
    r.checkNo = checkNo
    r.Set("check_no", checkNo)
    return nil
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetCheckNo() string {
    return r.checkNo
}

func (r *AlibabaAlihealthExaminationReserveReportNofifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveReportNofifyRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetType() string {
    return r.type
}

func (r *AlibabaAlihealthExaminationReserveReportNofifyRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaAlihealthExaminationReserveReportNofifyRequest) GetContent() string {
    return r.content
}

