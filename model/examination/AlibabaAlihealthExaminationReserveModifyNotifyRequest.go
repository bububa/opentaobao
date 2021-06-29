package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知改期结果 APIRequest
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功
*/
type AlibabaAlihealthExaminationReserveModifyNotifyRequest struct {
    model.Params

    // 旧的预约日期
    oldReserveDate   string 

    // 套餐编码
    packageCode   string 

    // 健康预约凭证
    reserveNumber   string 

    // 新的预约日期
    newReserveDate   string 

    // 服务商预约凭证
    uniqReserveCode   string 

    // 商品编码
    goodsCode   string 

    // 门店编码
    storeCode   string 

    // true:同意修改；false:拒绝修改
    pass   bool 

    // 拒绝修改的时候需要传递拒绝原因
    reason   string 

    // 新的预约时间段开始时间
    newReserveTimeStart   string 

    // 新的预约时间段结束时间
    newReserveTimeEnd   string 

}

func NewAlibabaAlihealthExaminationReserveModifyNotifyRequest() *AlibabaAlihealthExaminationReserveModifyNotifyRequest{
    return &AlibabaAlihealthExaminationReserveModifyNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.modify.notify"
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetOldReserveDate(oldReserveDate string) error {
    r.oldReserveDate = oldReserveDate
    r.Set("old_reserve_date", oldReserveDate)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetOldReserveDate() string {
    return r.oldReserveDate
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetPackageCode() string {
    return r.packageCode
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetNewReserveDate(newReserveDate string) error {
    r.newReserveDate = newReserveDate
    r.Set("new_reserve_date", newReserveDate)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetNewReserveDate() string {
    return r.newReserveDate
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetGoodsCode(goodsCode string) error {
    r.goodsCode = goodsCode
    r.Set("goods_code", goodsCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetGoodsCode() string {
    return r.goodsCode
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetPass(pass bool) error {
    r.pass = pass
    r.Set("pass", pass)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetPass() bool {
    return r.pass
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetNewReserveTimeStart(newReserveTimeStart string) error {
    r.newReserveTimeStart = newReserveTimeStart
    r.Set("new_reserve_time_start", newReserveTimeStart)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetNewReserveTimeStart() string {
    return r.newReserveTimeStart
}

func (r *AlibabaAlihealthExaminationReserveModifyNotifyRequest) SetNewReserveTimeEnd(newReserveTimeEnd string) error {
    r.newReserveTimeEnd = newReserveTimeEnd
    r.Set("new_reserve_time_end", newReserveTimeEnd)
    return nil
}

func (r AlibabaAlihealthExaminationReserveModifyNotifyRequest) GetNewReserveTimeEnd() string {
    return r.newReserveTimeEnd
}

