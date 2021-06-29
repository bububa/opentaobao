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
    _merchantCode   string
    // 阿里健康预约唯一标识
    _reserveNumber   string
    // 预约时间
    _reserveDate   string
    // 体检套餐编码
    _packageCode   string
    // 店铺ID
    _storeId   string
    // 体检机构预约唯一标识码
    _uniqReserveCode   string
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
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetMerchantCode() string {
    return r._merchantCode
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// ReserveDate Setter
// 预约时间
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetReserveDate() string {
    return r._reserveDate
}
// PackageCode Setter
// 体检套餐编码
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetPackageCode(_packageCode string) error {
    r._packageCode = _packageCode
    r.Set("package_code", _packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetPackageCode() string {
    return r._packageCode
}
// StoreId Setter
// 店铺ID
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetStoreId() string {
    return r._storeId
}
// UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveCancelRequest) SetUniqReserveCode(_uniqReserveCode string) error {
    r._uniqReserveCode = _uniqReserveCode
    r.Set("uniq_reserve_code", _uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveCancelRequest) GetUniqReserveCode() string {
    return r._uniqReserveCode
}
