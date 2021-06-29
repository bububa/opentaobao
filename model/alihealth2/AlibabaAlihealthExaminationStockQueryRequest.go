package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检套餐库存查询 APIRequest
alibaba.alihealth.examination.stock.query

体检机构对接_体检套餐库存查询
*/
type AlibabaAlihealthExaminationStockQueryRequest struct {
    model.Params

    // 商户唯一码
    merchantCode   string 

    // 分院ID
    hospitalId   string 

    // 套餐id
    packageId   string 

    // 开始日期
    timeFrom   string 

    // 结束日期
    timeTo   string 

}

func NewAlibabaAlihealthExaminationStockQueryRequest() *AlibabaAlihealthExaminationStockQueryRequest{
    return &AlibabaAlihealthExaminationStockQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.stock.query"
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationStockQueryRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaAlihealthExaminationStockQueryRequest) SetHospitalId(hospitalId string) error {
    r.hospitalId = hospitalId
    r.Set("hospital_id", hospitalId)
    return nil
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetHospitalId() string {
    return r.hospitalId
}

func (r *AlibabaAlihealthExaminationStockQueryRequest) SetPackageId(packageId string) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetPackageId() string {
    return r.packageId
}

func (r *AlibabaAlihealthExaminationStockQueryRequest) SetTimeFrom(timeFrom string) error {
    r.timeFrom = timeFrom
    r.Set("time_from", timeFrom)
    return nil
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetTimeFrom() string {
    return r.timeFrom
}

func (r *AlibabaAlihealthExaminationStockQueryRequest) SetTimeTo(timeTo string) error {
    r.timeTo = timeTo
    r.Set("time_to", timeTo)
    return nil
}

func (r AlibabaAlihealthExaminationStockQueryRequest) GetTimeTo() string {
    return r.timeTo
}

