package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV调TOP主动发起改期信息 API请求
alibaba.alihealth.examination.reserve.isv.modify

体检机构对接_ISV发起体检套餐改期
*/
type AlibabaAlihealthExaminationReserveIsvModifyRequest struct {
    model.Params
    // 阿里健康预约唯一标识
    _reserveNumber   string
    // 体检机构预约唯一标识码
    _uniqReserveCode   string
    // 修改后预约预约日期，格式：yyyy-MM-dd
    _reserveDate   string
    // 修改后预约时间段开始时间 HH:mm:ss
    _reserveTimeStart   string
    // 修改后预约时间段结束时间 HH:mm:ss
    _reserveTimeEnd   string
}

// 初始化AlibabaAlihealthExaminationReserveIsvModifyRequest对象
func NewAlibabaAlihealthExaminationReserveIsvModifyRequest() *AlibabaAlihealthExaminationReserveIsvModifyRequest{
    return &AlibabaAlihealthExaminationReserveIsvModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.isv.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetUniqReserveCode(_uniqReserveCode string) error {
    r._uniqReserveCode = _uniqReserveCode
    r.Set("uniq_reserve_code", _uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetUniqReserveCode() string {
    return r._uniqReserveCode
}
// ReserveDate Setter
// 修改后预约预约日期，格式：yyyy-MM-dd
func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveDate() string {
    return r._reserveDate
}
// ReserveTimeStart Setter
// 修改后预约时间段开始时间 HH:mm:ss
func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveTimeStart(_reserveTimeStart string) error {
    r._reserveTimeStart = _reserveTimeStart
    r.Set("reserve_time_start", _reserveTimeStart)
    return nil
}

// ReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveTimeStart() string {
    return r._reserveTimeStart
}
// ReserveTimeEnd Setter
// 修改后预约时间段结束时间 HH:mm:ss
func (r *AlibabaAlihealthExaminationReserveIsvModifyRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
    r._reserveTimeEnd = _reserveTimeEnd
    r.Set("reserve_time_end", _reserveTimeEnd)
    return nil
}

// ReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveIsvModifyRequest) GetReserveTimeEnd() string {
    return r._reserveTimeEnd
}
