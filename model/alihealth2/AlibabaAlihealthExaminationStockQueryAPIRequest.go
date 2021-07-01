package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检套餐库存查询 API请求
alibaba.alihealth.examination.stock.query

体检机构对接_体检套餐库存查询
*/
type AlibabaAlihealthExaminationStockQueryAPIRequest struct {
    model.Params
    // 商户唯一码
    _merchantCode   string
    // 分院ID
    _hospitalId   string
    // 套餐id
    _packageId   string
    // 开始日期
    _timeFrom   string
    // 结束日期
    _timeTo   string
}

// 初始化AlibabaAlihealthExaminationStockQueryAPIRequest对象
func NewAlibabaAlihealthExaminationStockQueryRequest() *AlibabaAlihealthExaminationStockQueryAPIRequest{
    return &AlibabaAlihealthExaminationStockQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.stock.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationStockQueryAPIRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetMerchantCode() string {
    return r._merchantCode
}
// HospitalId Setter
// 分院ID
func (r *AlibabaAlihealthExaminationStockQueryAPIRequest) SetHospitalId(_hospitalId string) error {
    r._hospitalId = _hospitalId
    r.Set("hospital_id", _hospitalId)
    return nil
}

// HospitalId Getter
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetHospitalId() string {
    return r._hospitalId
}
// PackageId Setter
// 套餐id
func (r *AlibabaAlihealthExaminationStockQueryAPIRequest) SetPackageId(_packageId string) error {
    r._packageId = _packageId
    r.Set("package_id", _packageId)
    return nil
}

// PackageId Getter
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetPackageId() string {
    return r._packageId
}
// TimeFrom Setter
// 开始日期
func (r *AlibabaAlihealthExaminationStockQueryAPIRequest) SetTimeFrom(_timeFrom string) error {
    r._timeFrom = _timeFrom
    r.Set("time_from", _timeFrom)
    return nil
}

// TimeFrom Getter
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetTimeFrom() string {
    return r._timeFrom
}
// TimeTo Setter
// 结束日期
func (r *AlibabaAlihealthExaminationStockQueryAPIRequest) SetTimeTo(_timeTo string) error {
    r._timeTo = _timeTo
    r.Set("time_to", _timeTo)
    return nil
}

// TimeTo Getter
func (r AlibabaAlihealthExaminationStockQueryAPIRequest) GetTimeTo() string {
    return r._timeTo
}
