package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV调TOP主动发起改期信息 APIRequest
alibaba.alihealth.examination.reserve.isv.modify

体检机构对接_ISV发起体检套餐改期
*/
type AlibabaAlihealthExaminationReserveIsvModifyRequest struct {
    model.Params

    // 阿里健康预约唯一标识
    reserveNumber   string 

    // 体检机构预约唯一标识码
    uniqReserveCode   string 

    // 修改后预约预约日期，格式：yyyy-MM-dd
    reserveDate   string 

    // 修改后预约时间段开始时间 HH:mm:ss
    reserveTimeStart   string 

    // 修改后预约时间段结束时间 HH:mm:ss
    reserveTimeEnd   string 

}

func NewAlibabaAlihealthExaminationReserveIsvModifyRequest() *AlibabaAlihealthExaminationReserveIsvModifyRequest{
    return &AlibabaAlihealthExaminationReserveIsvModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.isv.modify"
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveTimeStart(reserveTimeStart string) error {
    r.reserveTimeStart = reserveTimeStart
    r.Set("reserve_time_start", reserveTimeStart)
    return nil
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveTimeStart() string {
    return r.reserveTimeStart
}

func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveTimeEnd(reserveTimeEnd string) error {
    r.reserveTimeEnd = reserveTimeEnd
    r.Set("reserve_time_end", reserveTimeEnd)
    return nil
}

func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveTimeEnd() string {
    return r.reserveTimeEnd
}

