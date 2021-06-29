package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_预约取消 API请求
alibaba.alihealth.examination.reserve.cancel

体检机构对接_体检取消
*/
type AlibabaAlihealthExaminationReserveCancelRequest struct {
    model.Params
    // 商户唯一码
    merchantCode   string
    // 阿里健康预约唯一标识
    reserveNumber   string
    // 预约时间
    reserveDate   string
    // 体检套餐编码
    packageCode   string
    // 店铺ID
    storeId   string
    // 体检机构预约唯一标识码
    uniqReserveCode   string
}

// 初始化AlibabaAlihealthExaminationReserveCancelRequest对象
func NewAlibabaAlihealthExaminationReserveCancelRequest() *AlibabaAlihealthExaminationReserveCancelRequest{
    return &AlibabaAlihealthExaminationReserveCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetMerchantCode() string {
    return r.merchantCode
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetReserveNumber() string {
    return r.reserveNumber
}
// ReserveDate Setter
// 预约时间
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetReserveDate() string {
    return r.reserveDate
}
// PackageCode Setter
// 体检套餐编码
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetPackageCode() string {
    return r.packageCode
}
// StoreId Setter
// 店铺ID
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetStoreId() string {
    return r.storeId
}
// UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}
